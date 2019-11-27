package main

import (
	m "Misc"
)

// Initialization of ServerSide Implementation
func main() {
	go InitInput()
	InitConn(m.ServerConfiguration.NetType, m.ServerConfiguration.Address, m.ServerConfiguration.Port)
}
