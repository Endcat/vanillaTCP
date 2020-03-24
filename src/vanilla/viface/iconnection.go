package viface

import "net"

// define abstract layer of connection module

type IConnection interface {
	// start connection, ready to work
	Start()
	// stop connection, terminate
	Stop()
	// get current connection's socket conn
	GetTCPConnection() *net.TCPConn
	// get current connection's conn id
	GetConnID() uint32
	// get remote client addr
	RemoteAddr() net.Addr
	// send data to remote client
	SendMsg(msgId uint32, data []byte) error
}

// define function to handle connection service
type HandleFunc func(*net.TCPConn, []byte, int) error
