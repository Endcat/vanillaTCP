package vnet

import (
	"fmt"
	"strconv"
	"vanilla/viface"
)

type MsgHandle struct {
	// store handler for every msgid
	Apis map[uint32] viface.IRouter
}

// init/create msghandle
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{Apis:make(map[uint32] viface.IRouter)}
}

// call router message method
func (mh* MsgHandle) DoMsgHandler(request viface.IRequest){
	// get msgID from request
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("[Warning] api msgID = ",request.GetMsgID(), " is not found!")
	}
	// call router by msgID
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// add router message method
func (mh* MsgHandle) AddRouter(msgID uint32, router viface.IRouter) {
	// if already existed
	if _, ok := mh.Apis[msgID]; ok {
		panic("Repeat api, msgID = "+strconv.Itoa(int(msgID)))
	}
	// binding msg & api
	mh.Apis[msgID] = router
	fmt.Println("[MsgHandler] Add api MsgID = ",msgID, " Success!")
}