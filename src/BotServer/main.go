package main

import (
	m "Misc"
)

// Initialization of ServerSide Implemenation
func main() {
	go InitInput()
	InitConn(m.ServerConfiguration.NetType, m.ServerConfiguration.Address, m.ServerConfiguration.Port)
}
