package ziface

// IMessage 用于封装消息
type IMessage interface {

	// GetMsgId 获取消息ID
	GetMsgId() uint32

	GetDataLen() uint32

	GetData() []byte

	SetMsgId(id uint32)

	SetData(data []byte)

	SetDataLen(len uint32)
}
