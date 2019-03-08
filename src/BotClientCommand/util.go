package botclientcommand

import "net"

// SendToServer sends a message to server
func sendToServer(c net.Conn, msg string) (err error) {
	_, err = c.Write([]byte(msg))
	return
}
