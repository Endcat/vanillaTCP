package vnet
// server implement

import (
	"fmt"
	"net"
	"vanilla/viface"
)

type Server struct {
	// define server properties
	Name string
	IPVersion string
	IP string
	Port int
}

// define server methods (implement)
func (s *Server) Start() {
	// get server addr
	fmt.Printf("[start] Server Listener %s:%d\n",s.IP,s.Port)

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

		// waiting connection
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("[Error] Catch incoming call error ",err)
				continue
			}

			// connection established -> echo serving in maxlength 512
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("[Error] Catch buffer error ",err)
						continue
					}

					fmt.Printf("[Start] Receive client buffer %s, cnt %d\n",buf,cnt)
					// echo function
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("[Error] Catch write buffer error ",err)
						continue
					}
				}
			}()
		}
	}()
}
func (s *Server) Stop() {

}
func (s *Server) Serve() {
	s.Start()

	select {

	}
}

// server initiate
func NewServer(name string) viface.IServer{
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}