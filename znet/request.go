package znet

import "zinx-copy/ziface"

type Request struct {

	// 已经可客户端建立好的链接
	conn ziface.IConnection

	// 客户端请求的数据
	msg ziface.IMessage
}

// GetConnection 得到当前链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 得到当前数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 得到当前请求消息ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
