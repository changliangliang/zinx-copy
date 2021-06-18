package znet

import (
	"fmt"
	"zinx-copy/ziface"
)

// 消息处理模块实现

type MsgHandle struct {

	// 每个msgID对应的处理方法
	Apis map[uint32]ziface.IRouter
}

// NewMsgHandle 创建
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
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
