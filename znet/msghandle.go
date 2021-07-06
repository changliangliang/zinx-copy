package znet

import (
	"fmt"
	"zinx-copy/utils"
	"zinx-copy/ziface"
)

// 消息处理模块实现

type MsgHandle struct {

	// 每个msgID对应的处理方法
	Apis map[uint32]ziface.IRouter

	// 消息队列
	TaskQueue []chan ziface.IRequest

	// Worker池中goroutine数量
	WorkerPoolSize uint32
}

// NewMsgHandle 创建
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		return
	}
	m.Apis[msgId] = router
}

func (m *MsgHandle) DoMsgHandle(request ziface.IRequest) {
	if router, ok := m.Apis[request.GetMsgID()]; ok {
		router.PreHandle(request)
		router.Handle(request)
		router.PostHandle(request)
		return
	}
	fmt.Println("")

}

// StartWorkerPool 启动一个Worker工作池
func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {

		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

// StartOneWorker 启动一个工作流程
func (m *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {

	fmt.Println("Worker ID=", workerID, "is starting...")

	for request := range taskQueue {
		m.DoMsgHandle(request)
	}

}

// SendMsgToTaskQueue 将请求发送到指定worker
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {

	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(),
		"request MsgID=", request.GetMsgID(), "to WorkerID=", workerID)
	m.TaskQueue[workerID] <- request

}
