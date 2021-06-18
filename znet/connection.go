package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-copy/ziface"
)

// Connection 链接
type Connection struct {

	// 当前连接的socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChan chan bool

	MsgHandle ziface.IMsgHandle
}

// StartReader 链接读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
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

		// 处理请求
		go func() {

			// 从路由中找到对应的router
			c.MsgHandle.DoMsgHandle(req)

		}()

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
	if _, err := c.GetTCPConnection().Write(msg); err != nil {
		fmt.Println("Write msg id", msgID, "error:", err)
		return err
	}

	return nil

}

// Start 启动链接
func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID=", c.ConnID)
	// 启动从当前链接读数据的业务
	go c.StartReader()
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
	close(c.ExitChan)
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

func (c *Connection) Send(data []byte) error {
	panic("implement me")
}

// NewConnection 构造Connection
func NewConnection(conn *net.TCPConn, connID uint32, msgHandle ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		MsgHandle: msgHandle,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
