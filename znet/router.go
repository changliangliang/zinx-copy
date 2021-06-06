package znet

import "zinx-copy/ziface"

// BaseRouter 实现router时可以嵌入这个router基类
type BaseRouter struct{}

func (r *BaseRouter) PreHandle(request ziface.IRequest) {
}

func (r *BaseRouter) Handle(request ziface.IRequest) {
}

func (r *BaseRouter) PostHandle(request ziface.IRequest) {
}
