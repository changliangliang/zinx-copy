package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx-copy/ziface"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()

	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Println("Add conn to connManager, connID=", conn.GetConnID())

}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())
	fmt.Println("Remove conn from connManager, connID=", conn.GetConnID())

}

func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	}
	return nil, errors.New("conn not exist")

}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
	fmt.Println("Clear connections")

}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
