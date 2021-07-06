package znet

import (
	"fmt"
	"net"
	"zinx-copy/utils"
	"zinx-copy/ziface"
)

// Server IServer接口的实现
type Server struct {

	// 服务器名称
	Name string

	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的ip
	IP string

	// 服务器监听的端口
	Port int

	// router
	MsgHandle ziface.IMsgHandle

	// 链接管理器
	ConnManager ziface.IConnManager
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandle:   NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
}

// Start 启动服务器方法
func (s *Server) Start() {
	fmt.Printf(
		"[zinx] Server Name: %s, Listenner at IP:%s, Port:%d, is starting\n",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)
	fmt.Printf(
		"[zinx] Version %s, MaxConn %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize,
	)

	// 1、 获取TCP的Addr
	go func() {

		// 开启消息队列和worker工作池
		s.MsgHandle.StartWorkerPool()

		TCPaddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}
		// 监听服务地址
		TCPListener, err := net.ListenTCP(s.IPVersion, TCPaddr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err:", err)
			return
		}
		fmt.Println("start Zinx server succ, ", s.Name, " succ, Listening..")

		var cid uint32
		cid = 0
		for {
			conn, err := TCPListener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}

			// 判断是否超过最大连接数
			if s.ConnManager.Len() > utils.GlobalObject.MaxConn {
				// TODO 给客户端一个信息
				_ = conn.Close()
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandle)

			go dealConn.Start()
			cid++
		}
	}()

}

// Stop 关闭服务器方法
func (s *Server) Stop() {

	// 将连接资源回收
	s.ConnManager.ClearConn()
}

// Serve 运行服务器
func (s *Server) Serve() {
	s.Start()

	select {}

}

func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandle.AddRouter(msgID, router)
	fmt.Println("Add Router Succ!! ")

}
