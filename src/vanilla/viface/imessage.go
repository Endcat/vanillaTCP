package viface

// parse request data in Message

type IMessage interface {
	// get message id
	GetMsgId() uint32
	// get message length
	GetMsgLen() uint32
	// get message data
	GetData() []byte

	// set message id
	SetMsgId(uint32)
	// set message length
	SetDataLen(uint32)
	// set message data
	SetData([]byte)
}
