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

}

// Handle 处理请求
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	// 点读取客户端数据,再回写

	fmt.Println("recv from client: msgID=", request.GetMsgID(),
		"，data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
}

func server() {

	server := znet.NewServer("[zinx_v_0.3]")

	server.AddRouter(&PingRouter{})

	server.Serve()

}
