package viface
// abstract layer
// define server interface

type IServer interface {
	// basic functions: start/stop/serve
	Start()
	Stop()
	Serve()
	// register router method
	AddRouter(msgID uint32, router IRouter)
	// get current server connection manager
	GetConnMgr() IConnManager
	// register hook function
	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))
	// call hook function
	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)
}
