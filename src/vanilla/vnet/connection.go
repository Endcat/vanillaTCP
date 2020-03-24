package vnet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"vanilla/viface"
)

// connection module

type Connection struct {
	// current connection socket
	Conn *net.TCPConn

	// connection ID
	ConnID uint32

	// current connection status
	isClosed bool

	//// current connection binding service API
	//handleAPI viface.HandleFunc

	// exit channel
	ExitChan chan bool

	// current connection router
	Router viface.IRouter
}

// initiate connection method
func NewConnection(conn *net.TCPConn, connID uint32, router viface.IRouter) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		Router:    router,
		//handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// connection reader service
func (c *Connection) StartReader() {
	fmt.Println("[Start] Reader goroutine launching...")
	defer fmt.Println("[Stop] connID = ",c.ConnID, " Reader terminated, remote addr = ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("[Error] Catch receive buffer error ", err)
		//	continue
		//}

		// create pack/unpack object
		dp := NewDataPack()

		// read client msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("[Error] Catch read message head error ",err)
			break
		}

		// unpack
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("[Error] Catch unpack error ",err)
			break
		}

		// read data
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("[Error] Catch read message data error ",err)
				break
			}
		}
		msg.SetData(data)

		// get current request data
		req := Request{
			conn:c,
			msg:msg,
		}

		// execute router methods
		go func(request viface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		//// call current connection's handleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("[Error] Catch handle error ",err)
		//	break
		//}
	}
}


// start connection, ready to work
func (c *Connection) Start() {
	fmt.Println("[Start] Connection start.. ConnID = ", c.ConnID)
	// launch reader service
	go c.StartReader()
}
// stop connection, terminate
func (c *Connection) Stop() {
	fmt.Println("[Stop] Connection stop.. ConnID = ",c.ConnID)

	// if already closed
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// close socket connection
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("[Error] Catch close connection error ",err)
		return
	}

	// gc
	close(c.ExitChan)

}
// get current connection's socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
// get current connection's conn id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
// get remote client addr
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
// send data to remote client
//func (c *Connection) Send(data []byte) error {
//	return nil
//}

// send message
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when sending message")
	}

	// pack data
	dp := NewDataPack()

	// msgdatalen | msgid | data
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("[Error] Catch pack error, msg id = ",msgId)
		return errors.New("Pack message error ")
	}

	// send data to client
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("[Error] Catch write message error, msg id = ",msgId," error: ",err)
		return errors.New("connection write data error")
	}

	return nil
}