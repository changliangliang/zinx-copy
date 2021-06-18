package ziface

// 消息管理抽象层

type IMsgHandle interface {
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)
	// DoMsgHandle 调度执行消息对应的处理器
	DoMsgHandle(request IRequest)
}
