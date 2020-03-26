package vnet

import (
	"errors"
	"fmt"
	"sync"
	"vanilla/viface"
)

// connection management module

type ConnManager struct {
	connections map[uint32]viface.IConnection
	connLock    sync.RWMutex
}

// create conn manager
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]viface.IConnection),
		connLock:    sync.RWMutex{},
	}
}

// add conn
func (connMgr *ConnManager) Add(conn viface.IConnection) {
	// protect map
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// add conn to connManager
	connMgr.connections[conn.GetConnID()] = conn

	fmt.Println("[ConnMgr] connID = ", conn.GetConnID(), " Add conn to connMgr successfully: conn num = ", connMgr.Len())
}

// delete conn
func (connMgr *ConnManager) Remove(conn viface.IConnection) {
	// protect map
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// delete conn to connManager
	delete(connMgr.connections, conn.GetConnID())

	fmt.Println("[ConnMgr] connID = ", conn.GetConnID(), " Remove conn from connMgr successfully: conn num = ", connMgr.Len())
}

// get conn by connid
func (connMgr *ConnManager) Get(connID uint32) (viface.IConnection, error) {
	// protect map
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// current amount of conns
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// terminate all conns
func (connMgr *ConnManager) ClearConn() {
	// protect map
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	for connID, conn := range connMgr.connections {
		conn.Stop()
		delete(connMgr.connections, connID)
	}

	fmt.Println("[ConnMgr] Clear all connections successfully! conn num = ", connMgr.Len())
}
