package vnet
// server implement

import (
	"fmt"
	"net"
	"vanilla/utils"
	"vanilla/viface"
)

type Server struct {
	// define server properties
	Name string
	IPVersion string
	IP string
	Port int
	// current registered router
	Router viface.IRouter
}

//// define current client binding handleAPI
//func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
//	// echo service
//	fmt.Println("[Handle] Echo service... ")
//	if _,err := conn.Write(data[:cnt]); err != nil {
//		fmt.Println("[Error] Echo buffer error ",err)
//		return errors.New("CallBackToClient error")
//	}
//	return nil
//}

// define server methods (implement)
func (s *Server) Start() {
	// print info
	fmt.Printf("[Vanilla] Server name: %s, Listener at IP: %s, Port: %d is starting\n",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)
	fmt.Printf("[Vanilla] Version %s, MaxConn: %d, MaxPacketSize: %d\n]",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err != nil {
			fmt.Println("[Error] Catch TCP address error: ",err)
			return
		}

		// listen server addr
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil{
			fmt.Println("[Error] Catch TCP listener error: ",err)
			return
		}

		// create listener successfully, listening
		fmt.Println("[Start] Start Vanilla server successfully, Name = ",s.Name, "Listening...")
		var cid uint32
		cid = 0

		// waiting connection
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("[Error] Catch incoming call error ",err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
			//// connection established -> echo serving in maxlength 512
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("[Error] Catch buffer error ",err)
			//			continue
			//		}
			//
			//		fmt.Printf("[Start] Receive client buffer %s, cnt %d\n",buf,cnt)
			//		// echo function
			//		if _, err := conn.Write(buf[:cnt]); err != nil {
			//			fmt.Println("[Error] Catch write buffer error ",err)
			//			continue
			//		}
			//	}
			//}()
		}
	}()
}
func (s *Server) Stop() {

}
func (s *Server) Serve() {
	s.Start()

	select {}
}

// register router methods for current server
func (s *Server) AddRouter(router viface.IRouter) {
	s.Router = router
	fmt.Println("[Router] Successfully added router!")
}

// server initiate
func NewServer(name string) viface.IServer{
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:nil,
	}
	return s
}