package ziface

// IServer 定义一个服务器接口
type IServer interface {

	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Serve 运行服务器
	Serve()

	//AddRouter 添加router
	AddRouter(msgID uint32, router IRouter)

	//GetConnManager 获得链接管理器
	GetConnManager() IConnManager

	// SetOnConnStart 注册OnConnStart
	SetOnConnStart(func(conn IConnection))

	// SetOnConnStop 注册OnConnStop
	SetOnConnStop(func(conn IConnection))

	// CallOnConnStart 调用OnConnStart
	CallOnConnStart(conn IConnection)

	// CallOnConnStop 调用OnConnStop
	CallOnConnStop(conn IConnection)
}
