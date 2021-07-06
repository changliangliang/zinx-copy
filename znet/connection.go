package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-copy/utils"
	"zinx-copy/ziface"
)

// Connection 链接
type Connection struct {

	// 当前Connection隶属的Server
	Server ziface.IServer

	// 当前连接的socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChan chan bool

	// 用于读写goroutine之间的通信
	msgChan chan []byte

	MsgHandle ziface.IMsgHandle
}

// StartReader 链接读业务方法
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 创建一个拆包解包对象
		dp := NewDataPack()

		// 读取客户端 Msg Head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error:", err)
			break
		}

		// 拆包得到msgID和 msgDataLen 放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:", err)
			break
		}

		// 根据datalen 再次读取data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err = io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error:", err)
				break
			}
			msg.SetData(data)
		}

		// 得到Request
		req := &Request{
			conn: c,
			msg:  msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandle.SendMsgToTaskQueue(req)

		} else {
			// 处理请求
			go c.MsgHandle.DoMsgHandle(req)

		}

	}
}

// SendMsg 发送消息方法
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	// 将data进行封包
	dp := NewDataPack()
	msg, err := dp.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println("Pack msg error:", err)
		return err
	}
	c.msgChan <- msg

	return nil

}

// Start 启动链接
func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID=", c.ConnID)

	// 启动读业务
	go c.StartReader()

	// 启动写业务
	go c.StartWriter()
}

// Stop 关闭链接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID=", c.ConnID)

	// 当前链接已经关闭
	if c.isClosed == true {
		fmt.Println("Conn has closed")
		return
	}

	c.isClosed = false

	// 回收资源
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Conn close error:", err)
		return
	}

	c.ExitChan <- true

	c.Server.GetConnManager().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)
	fmt.Println("Conn has closed")
}

// GetTCPConnection 获取当前链接绑定的socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端TCP状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// StartWriter 写消息的goroutine, 专门给客户端发送消息
func (c *Connection) StartWriter() {
	fmt.Println("[writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[con Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error", err)
				return
			}
		case <-c.ExitChan:
			// 代表reader退出
			return

		}
	}

}

// NewConnection 构造Connection
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandle ziface.IMsgHandle) *Connection {
	c := &Connection{
		Server:    server,
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		MsgHandle: msgHandle,
		msgChan:   make(chan []byte),
		ExitChan:  make(chan bool, 1),
	}

	c.Server.GetConnManager().Add(c)

	return c
}
