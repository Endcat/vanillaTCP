package main

import (
	"fmt"
	"vanilla/viface"
	"vanilla/vnet"
)

// router test
type HelloVanillaRouter struct {
	vnet.BaseRouter
}

func (this *HelloVanillaRouter) Handle(request viface.IRequest) {
	fmt.Println("[Server] Call HelloVanillaRouter handle...")
	fmt.Println("[Server] Recv from client: msgID = ",request.GetMsgID(), ", data = ",string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("zxzxzxzxzxzx"))
	if err != nil {
		fmt.Println(err)
	}
}

type PingRouter struct {
	vnet.BaseRouter
}

func (this *PingRouter) Handle(request viface.IRequest) {
	fmt.Println("[Server] Call HelloVanillaRouter handle...")
	fmt.Println("[Server] Recv from client: msgID = ",request.GetMsgID(), ", data = ",string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("zxzxzxzxzxzx"))
	if err != nil {
		fmt.Println(err)
	}
}

// hook function
func DoConnectionBegin(conn viface.IConnection) {
	fmt.Println("[Server] DoConnectionBegin is called")
	if err := conn.SendMsg(202, []byte("DoConnectionBegin is called")); err != nil {
		fmt.Println(err)
	}

	// set property
	fmt.Println("[Server] Setting properties ...")
	conn.SetProperty("Name", "vanilla")
}

func DoConnectionLost(conn viface.IConnection) {
	fmt.Println("[Server] DoConnectionLost is called")
	fmt.Println("[Server] conn ID = ", conn.GetConnID(), " terminated")

	// get property
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("[Server] Get property Name = ", name)
	}
}

func main()  {
	// create server handler with Vanilla api
	s := vnet.NewServer("[Test 0x01]")

	// register hook function
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)


	// add router
	s.AddRouter(0, &HelloVanillaRouter{})
	s.AddRouter(1, &PingRouter{})

	// launch server
	s.Serve()
}

// useless code fragment


// prehandle test
//func (this *HelloVanillaRouter) PreHandle(request viface.IRequest) {
//	fmt.Println("[Server] Call router PreHandle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
//	if err != nil {
//		fmt.Println("[Error] Catch before ping error")
//	}
//}
// handle test
//func (this *HelloVanillaRouter) Handle(request viface.IRequest) {
//	fmt.Println("[Server] Call router Handle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping..."))
//	if err != nil {
//		fmt.Println("[Error] Catch ping error")
//	}
//}
// posthandle test
//func (this *HelloVanillaRouter) PostHandle(request viface.IRequest) {
//	fmt.Println("[Server] Call router PostHandle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
//	if err != nil {
//		fmt.Println("[Error] Catch after ping error")
//	}
//}