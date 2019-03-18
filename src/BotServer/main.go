package main

import (
	m "Misc"
)

// Init ServerSide Implemenation
func main() {
	go InitInput()
	InitConn(m.ServerConfiguration.NetType, m.ServerConfiguration.Address, m.ServerConfiguration.Port)
}
