package ziface

// 封包拆包的数据

type IDataPack interface {

	// GetHeadLen 获取长度
	GetHeadLen() uint32
	// Pack 封包
	Pack(msg IMessage) ([]byte, error)
	// Unpack 拆包
	Unpack([]byte) (IMessage, error)
}
