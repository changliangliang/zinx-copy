package main

import (
	"fmt"
	"net"
	"time"
	"zinx-copy/znet"
)

func main() {

	go func() {
		time.Sleep(1 * time.Second)
		tcp, _ := net.Dial("tcp4", "127.0.0.1:8999")
		count := 1
		bytes := make([]byte, 512)

		for {
			_, _ = tcp.Write([]byte(fmt.Sprintf("[client]===>%d", count)))
			read, _ := tcp.Read(bytes)
			fmt.Println(string(bytes[0:read]))
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	server := znet.NewServer("zinx-copy")
	server.Serve()
}
