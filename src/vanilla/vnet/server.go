package vnet
// server implement

import "vanilla/viface"

type Server struct {
	// define server properties
	Name string
	IPVersion string
	IP string
	Port int
}

// define server methods (implement)
func (s *Server) Start() {
	
}
func (s *Server) Stop() {

}
func (s *Server) Serve() {
	s.Start()
}

// server initiate
func NewServer(name string) viface.IServer{
	s := &Server{
		Name:      name,
		IPVersion: "TCP4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}