package viface

// message handler interface

type IMsgHandle interface {
	// call router message method
	DoMsgHandler(request IRequest)
	// add router message method
	AddRouter(msgID uint32, router IRouter)
}