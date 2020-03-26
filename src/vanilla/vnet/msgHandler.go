package vnet

import (
	"fmt"
	"strconv"
	"vanilla/utils"
	"vanilla/viface"
)

type MsgHandle struct {
	// store handler for every msgid
	Apis map[uint32] viface.IRouter
	// task queue for worker
	TaskQueue []chan viface.IRequest
	// worker pool size
	WorkerPoolSize uint32
}

// init/create msghandle
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32] viface.IRouter),
		TaskQueue:      make([]chan viface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
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

// start worker pool (1 maximum)
func (mh* MsgHandle) StartWorkerPool() {
	// start worker individually, one worker one goroutine
	for i:=0; i < int(mh.WorkerPoolSize); i++ {
		// create channel for new worker
		mh.TaskQueue[i] = make(chan viface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// start worker
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// start worker
func (mh* MsgHandle) StartOneWorker(workerID int, taskQueue chan viface.IRequest) {
	fmt.Println("[Worker] Worker ID = ",workerID, " is started.")

	// waiting for task queue message
	for {
		select {
		case request := <- taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request viface.IRequest) {
	// equally distribution
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("[Worker] Add ConnID = ",request.GetConnection().GetConnID(),
		" request MsgID = ", request.GetMsgID(),
		" to workerID = ", workerID)

	// send message to task queue
	mh.TaskQueue[workerID] <- request
}