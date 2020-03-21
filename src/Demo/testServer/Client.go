package main

import (
	"fmt"
	"net"
	"time"
)

// simulate client

func main() {
	fmt.Println("[Client] Client start...")
	time.Sleep(1 * time.Second)

	// connect to server, get connection
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil{
		fmt.Println("[Error] Catch client start error ",err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello Vanilla!"))
		if err != nil{
			fmt.Println("[Error] Catch write conn error ",err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("[Error] Catch read buffer error ",err)
			return
		}

		fmt.Printf("[Client] Server echo: %s, cnt = %d\n",buf, cnt)

		time.Sleep(1 * time.Second)
	}
}