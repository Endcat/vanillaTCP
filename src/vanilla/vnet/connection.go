package vnet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"vanilla/utils"
	"vanilla/viface"
)

// connection module

type Connection struct {
	// current connected server
	TcpServer viface.IServer

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

	// read/write go routine
	msgChan chan []byte

	// current connection router
	//Router viface.IRouter
	// message handler
	MsgHandler viface.IMsgHandle
}

// initiate connection method
func NewConnection(server viface.IServer, conn *net.TCPConn, connID uint32, msgHandler viface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer: server,
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		msgChan: make(chan []byte),
		//Router:    router,
		//handleAPI: callbackAPI,
		MsgHandler: msgHandler,
		ExitChan:  make(chan bool, 1),
	}

	// add conn to connection manager
	c.TcpServer.GetConnMgr().Add(c)

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

		// if worker pool has started
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			// call message handler
			go c.MsgHandler.DoMsgHandler(&req)
		}


		// execute router methods
		//go func(request viface.IRequest) {
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)

		//// call current connection's handleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("[Error] Catch handle error ",err)
		//	break
		//}
	}
}

// write go routine
func (c *Connection) StartWriter() {
	fmt.Println("[Start] Writer goroutine launching...")
	defer fmt.Println("[Stop] connID = ",c.ConnID, " Writer terminated, remote addr = ", c.RemoteAddr().String())

	for {
		select {
		case data := <-c.msgChan:
			// write data to client
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("[Error] Catch send data error ",err)
				return
			}

		case <-c.ExitChan:
			// reader terminated, terminate writer
			return
		}

	}
}

// start connection, ready to work
func (c *Connection) Start() {
	fmt.Println("[Start] Connection start.. ConnID = ", c.ConnID)
	// launch reader service
	go c.StartReader()
	// launch writer service
	go c.StartWriter()

	// call hook function
	c.TcpServer.CallOnConnStart(c)
}
// stop connection, terminate
func (c *Connection) Stop() {
	fmt.Println("[Stop] Connection stop.. ConnID = ",c.ConnID)

	// if already closed
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// call hook function
	c.TcpServer.CallOnConnStop(c)

	// close socket connection
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("[Error] Catch close connection error ",err)
		return
	}

	// notice to terminate writer
	c.ExitChan <- true

	// remove connection from connMgr
	c.TcpServer.GetConnMgr().Remove(c)

	// gc
	close(c.ExitChan)
	close(c.msgChan)

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
	c.msgChan <- binaryMsg

	//if _, err := c.Conn.Write(binaryMsg); err != nil {
	//	fmt.Println("[Error] Catch write message error, msg id = ",msgId," error: ",err)
	//	return errors.New("connection write data error")
	//}

	return nil
}