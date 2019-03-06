package main

import (
	botcommand "BotCommand"
	misc "Misc"
	"net"
)

// InitConn initialize the connection
func InitConn(netType string, address string, port string) {
	myConn, err := net.Listen(netType, address+":"+port)
	if err != nil {
		misc.OutputLogger.Fatalln("Error Creating Server!", err)
	}
	misc.Info("Server Has Started")
	for {
		// Listen for an incoming connection.
		botClient, err := myConn.Accept()
		if err != nil {
			continue
		}
		// Handle connections in a new goroutine.
		go botHandler(botClient)
	}
}

func botHandler(c net.Conn) {

	for {
		buf := make([]byte, 1024)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}
		data := buf[0:nr]
		go botcommand.ReadBotMessage(c, string(data))
		// fmt.Printf("->: %s", string(data))
		// _, err = c.Write(data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
}
