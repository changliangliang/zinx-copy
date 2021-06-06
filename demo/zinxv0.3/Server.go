package main

import (
	"fmt"
	"zinx-copy/ziface"
	"zinx-copy/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("Call back before ping err:", err)
	}
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping..."))
	if err != nil {
		fmt.Println("Call back ping err:", err)
	}

}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
	if err != nil {
		fmt.Println("Call back after ping err:", err)
	}
}

func main() {

	server := znet.NewServer("[zinx_v_0.3]")

	server.AddRouter(&PingRouter{})

	server.Serve()

}
