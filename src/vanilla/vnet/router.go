package vnet

import "vanilla/viface"

type BaseRouter struct {}

// override methods individually
// prev hook method
func (br *BaseRouter) PreHandle(request viface.IRequest){}
// main hook method
func (br *BaseRouter) Handle(request viface.IRequest){}
// post hook method
func (br *BaseRouter) PostHandle(request viface.IRequest){}