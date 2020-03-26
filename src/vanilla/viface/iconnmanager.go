package viface

// connection management module

type IConnManager interface {
	// add conn
	Add(conn IConnection)
	// delete conn
	Remove(conn IConnection)
	// get conn by connid
	Get(connID uint32) (IConnection, error)
	// current amount of conns
	Len() int
	// terminate all conns
	ClearConn()
}