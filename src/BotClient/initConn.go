package main

import (
	botclientcommand "BotClientCommand"
	misc "Misc"
	"bufio"
	"net"
	"os"
)

//Conn ...
var Conn net.Conn

// initCLientConn initializes client connection
func initClientConn(netType string, addr string, port string) {
	Conn, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		misc.Err("cannot connect to server")
	}
	misc.Info("Client Connected")
	go handleClientInput(Conn)
	handleClientConn(Conn)
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	text, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		misc.Err("connot read from input", err)
	// 	}
	// 	handleClientInput(text)
	// }
}

func handleClientConn(c net.Conn) {
	_, err := c.Write([]byte(">setbotname " + botclientcommand.Getnick()))
	if err != nil {
		// return
	}
	for {
		buf := make([]byte, 500)
		nr, err := c.Read(buf)
		if err != nil {

		}
		data := string(buf[0:nr])
		if data == "Name Taken" {
			_, err = c.Write([]byte(">setbotname " + botclientcommand.Getnick()))
		}
		botclientcommand.HandleCommand(c, data)
	}
}

func handleClientInput(c net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			misc.Warning("cannot send a message", err)
			continue
		}
		c.Write([]byte(msg + "\n"))
	}
}
