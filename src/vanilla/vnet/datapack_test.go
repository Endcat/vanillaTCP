package vnet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// unit test for datapack

func TestDataPack(t *testing.T) {
	// simulate server
	// create tcp socket
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("[Error] Catch server listen err: ",err)
		return
	}

	// dealing service
	go func() {
		// read data & unpack
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[Error] Catch server accept error",err)
		}

		go func(conn net.Conn) {
			// dealing client request
			// unpack
			dp := NewDataPack()
			for {
				// read head
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("[Error] Catch read head error")
					break
				}

				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("[Error] Catch server unpack error ",err)
					return
				}
				if msgHead.GetMsgLen() > 0 {
					// read data
					msg := msgHead.(*Message)
					msg.Data = make([]byte, msg.GetMsgLen())

					// read by datalen
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("[Error] Catch server unpack error ",err)
						return
					}

					// read complete
					fmt.Println("[Recv] Recv MsgID: ", msg.Id, " | dataLen = ",msg.DataLen," | data = ", string(msg.Data))
				}
			}

		}(conn)
	}()

	// simulate client
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("[Error] Catch client dial error: ",err)
		return
	}

	// create data pack
	dp := NewDataPack()

	// packing process (2 examples)
	msg1 := &Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'v','a','n','i','l','l','a'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("[Error] Catch client pack msg1 error: ",err)
		return
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 6,
		Data:    []byte{'e','n','d','c','a','t'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("[Error] Catch client pack msg2 error: ",err)
		return
	}

	sendData1 = append(sendData1,sendData2...)
	if _, err := conn.Write(sendData1); err != nil {
		fmt.Println("[Error] Catch send pack msg error: ",err)
		return
	}

	select {}
}
