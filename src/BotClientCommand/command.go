package botclientcommand

import (
	"net"
)

type client struct {
	conn net.Conn
	name string
	room string
}

type room struct {
	members map[net.Addr]*client
}

var rooms map[string]*room

var usage = `
/nick <name>: get a name, or stay anonymous
/join <room>:

`