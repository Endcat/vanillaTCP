package main

import (
	"fmt"
	"vanilla/viface"
	"vanilla/vnet"
)

// router test
type PingRouter struct {
	vnet.BaseRouter
}

func (this *PingRouter) Handle(request viface.IRequest) {
	fmt.Println("[Server] Call router handle...")
	fmt.Println("[Server] Recv from client: msgID = ",request.GetMsgID(), ", data = ",string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("zxzxzxzxzxzx"))
	if err != nil {
		fmt.Println(err)
	}
}

func main()  {
	// create server handler with Vanilla api
	s := vnet.NewServer("[Test 0x01]")
	// add router
	s.AddRouter(&PingRouter{})
	// launch server
	s.Serve()
}

// useless code fragment


// prehandle test
//func (this *PingRouter) PreHandle(request viface.IRequest) {
//	fmt.Println("[Server] Call router PreHandle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
//	if err != nil {
//		fmt.Println("[Error] Catch before ping error")
//	}
//}
// handle test
//func (this *PingRouter) Handle(request viface.IRequest) {
//	fmt.Println("[Server] Call router Handle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping..."))
//	if err != nil {
//		fmt.Println("[Error] Catch ping error")
//	}
//}
// posthandle test
//func (this *PingRouter) PostHandle(request viface.IRequest) {
//	fmt.Println("[Server] Call router PostHandle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
//	if err != nil {
//		fmt.Println("[Error] Catch after ping error")
//	}
//}