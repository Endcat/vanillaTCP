package vnet

import "vanilla/viface"

type Request struct {
	// established connection with client
	conn viface.IConnection
	// client request data
	msg viface.IMessage
}

// get current connection
func (r *Request) GetConnection() viface.IConnection {
	return r.conn
}
// get request data
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}