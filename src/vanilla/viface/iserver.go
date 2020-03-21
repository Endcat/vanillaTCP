package viface
// abstract layer
// define server interface

type IServer interface {
	// basic functions: start/stop/serve
	Start()
	Stop()
	Serve()
}