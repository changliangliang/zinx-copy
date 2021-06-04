package ziface

import "net"

// 链接模块抽象层

type IConnection interface {

	// Start 启动链接
	Start()

	// Stop 停止链接
	Stop()

	// GetTCPConnection 获取绑定socket
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前链接id
	GetConnID() uint32

	// RemoteAddr 获取远程客户端TCP状态
	RemoteAddr() net.Addr

	// Send 发送数据
	Send([]byte) error
}

// HandleFunc 处理业务的方法
type HandleFunc func(conn *net.TCPConn, data []byte, len int) error
