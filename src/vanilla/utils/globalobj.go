package utils

import (
	"encoding/json"
	"io/ioutil"
	"vanilla/viface"
)

// global params

type GlobalObj struct {
	// server setting
	TcpServer viface.IServer	// current global server object
	Host string					// current listening ip
	TcpPort int					// current listening port
	Name string					// current server name
	// vanilla setting
	Version string				// current vanilla version
	MaxConn int					// current server maximum allowed connections
	MaxPackageSize uint32		// current vanilla maximum data package size

}

// define public global object
var GlobalObject *GlobalObj

// load params from vanilla.json
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/vanilla.json")
	if err != nil {panic(err)}

	// parse vanilla.json
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {panic(err)}
}

// initiation
func init() {
	// default
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Name:           "VanillaServerApp",
		Version:        "V0.1",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	// load from conf/vanilla.json
	GlobalObject.Reload()
}