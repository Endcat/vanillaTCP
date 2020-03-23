package vnet

import "vanilla/viface"

type Request struct {
	// established connection with client
	conn viface.IConnection
	// client request data
	data []byte
}

// get current connection
func (r *Request) GetConnection() viface.IConnection {
	return r.conn
}
// get request data
func (r *Request) GetData() []byte {
	return r.data
}