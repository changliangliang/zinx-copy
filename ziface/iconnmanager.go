package ziface

type IConnManager interface {

	// Add 添加链接
	Add(conn IConnection)

	// Remove 删除链接
	Remove(conn IConnection)

	// Get 根据id获得链接
	Get(connId uint32) (IConnection, error)

	// Len 链接总数
	Len() int

	// ClearConn 关闭所有链接
	ClearConn()
}
