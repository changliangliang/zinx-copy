package znet

import (
	"errors"
	"fmt"
	"net"
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
}

// CallBackToClient 当前回调是写死的，以后应该有用户自定义
func CallBackToClient(conn *net.TCPConn, data []byte, len int) error {
	fmt.Println("[Conn Handle] CallBackToClient...]")
	if _, err := conn.Write(data[0:len]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("CallBackToClient error")
	}
	return nil

}

func NewServer(name string) ziface.IServer {

	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

// Start 启动服务器方法
func (s Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port:%d, is starting\n", s.IP, s.Port)
	// 1、 获取TCP的Addr
	go func() {
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
			dealConn := NewConnection(conn, cid, CallBackToClient)
			go dealConn.Start()
			cid++
		}
	}()

}

// Stop 关闭服务器方法
func (s Server) Stop() {
	// TODO

}

// Serve 运行服务器
func (s Server) Serve() {
	s.Start()

	select {}

}