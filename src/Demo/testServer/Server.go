package main

import "vanilla/vnet"

func main()  {
	// create server handler with Vanilla api
	s := vnet.NewServer("[Test 0x01]")
	// launch server
	s.Serve()
}
