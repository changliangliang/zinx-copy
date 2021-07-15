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

	// SendMsg 发送数据
	SendMsg(msgID uint32, data []byte) error

	// SetProperty 设置链接属性
	SetProperty(key string, property interface{})

	// GetProperty 获取属性
	GetProperty(key string) (interface{}, error)

	// RemoveProperty 移除属性
	RemoveProperty(key string)
}

// HandleFunc 处理业务的方法
type HandleFunc func(conn *net.TCPConn, data []byte, len int) error
