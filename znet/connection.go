package znet

import (
	"fmt"
	"net"
	"zinx-copy/ziface"
)

// Connection 链接
type Connection struct {

	// 当前连接的socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	handleApi ziface.HandleFunc

	ExitChan chan bool
}

// StartReader 链接读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端数据到buf中
		buf := make([]byte, 512)
		read, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 调用当前链接绑定的HandleApi
		if err := c.handleApi(c.Conn, buf, read); err != nil {
			fmt.Println("ConnID", c.ConnID, "handle is error", err)
		}
	}
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
func NewConnection(conn *net.TCPConn, connID uint32, callBackApi ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleApi: callBackApi,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
