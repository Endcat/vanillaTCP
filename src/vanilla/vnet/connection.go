package vnet

import (
	"net"
	"vanilla/viface"
)

// connection module

type Connection struct {
	// current connection socket
	Conn *net.TCPConn

	// connection ID
	ConnID uint32

	// current connection status
	isClosed bool

	// current connection binding service API
	handleAPI viface.HandleFunc

	// exit channel
	ExitChan chan bool
}

// initiate connection method
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI viface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// start connection, ready to work
func (c *Connection) Start() {

}
// stop connection, terminate
func (c *Connection) Stop() {

}
// get current connection's socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
// get current connection's conn id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
// get remote client addr
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
// send data to remote client
func (c *Connection) Send(data []byte) error {

}