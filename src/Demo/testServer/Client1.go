package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"vanilla/vnet"
)

// simulate client

func main() {
	fmt.Println("[Client] Client1 start...")
	time.Sleep(1 * time.Second)

	// connect to server, get connection
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil{
		fmt.Println("[Error] Catch client start error ",err)
		return
	}

	for {
		dp := vnet.NewDataPack()
		binaryMsg, err := dp.Pack(vnet.NewMsgPackage(1, []byte("Vanilla client1 test")))
		if err != nil {
			fmt.Println("[Error] Catch pack error", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("[Error] Catch write error",err)
			return
		}

		// read data head
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("[Error] Catch read head error ",err)
			break
		}

		// unpack
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("[Error] Catch client unpack error ",err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			// read data
			msg := msgHead.(*vnet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("[Error] Catch read msg data error, ",err)
				return
			}

			fmt.Println("[Server] Recv server message: ID = ",msg.Id, ", len = ", msg.DataLen, ", data = ", string(msg.Data))
			
		}

		time.Sleep(1 * time.Second)
	}
}

// useless code fragment

//for {
//_, err := conn.Write([]byte("Hello Vanilla!"))
//if err != nil{
//fmt.Println("[Error] Catch write conn error ",err)
//return
//}
//
//buf := make([]byte, 512)
//cnt, err := conn.Read(buf)
//if err != nil {
//fmt.Println("[Error] Catch read buffer error ",err)
//return
//}
//
//fmt.Printf("[Client] Server echo: %s, cnt = %d\n",buf, cnt)
//
//time.Sleep(1 * time.Second)
//}