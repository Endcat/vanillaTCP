package viface

// abstract router interface

type IRouter interface {
	// prev hook method
	PreHandle(request IRequest)
	// main hook method
	Handle(request IRequest)
	// post hook method
	PostHandle(request IRequest)
}