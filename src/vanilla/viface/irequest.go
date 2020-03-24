package viface

// IRequest Interface
// pack client request connection info & request data

type IRequest interface {
	// get current connection
	GetConnection() IConnection
	// get request data
	GetData() []byte
	// get current request message id
	GetMsgID() uint32
}