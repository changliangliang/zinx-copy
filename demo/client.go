package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx-copy/znet"
)

func client() {
	fmt.Println("Client start......")

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error:", err)
		return
	}

	for {

		// 发送分包消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMessage(1, []byte("Zinxv0.5 client Test Message....")))
		if err != nil {
			fmt.Println("Pack error:", err)
			break
		}
		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("Write msg error：", err)
			break
		}

		// 服务器会返回给我们一个数据
		binaryHead := make([]byte, dp.GetHeadLen())

		if _, err = io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error:", err)
			break
		}

		msg, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("unpack error:", err)
			break
		}

		if msg.GetDataLen() > 0 {

			msg := msg.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {

				fmt.Println("read msgdata error:", err)
				break
			}

			fmt.Println("----> recv serer msg: id=", msg.GetMsgId(), "，len=", msg.GetDataLen(), ", data=", string(msg.GetData()))
		}

		time.Sleep(1 * time.Second)

	}
}
