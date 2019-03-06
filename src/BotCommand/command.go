package botcommand

import (
	misc "Misc"
	"data"
	"fmt"
	"net"
	"strings"
)

// ReadBotMessage Reads the message comming from a bot and make it ready to process
func ReadBotMessage(botConn net.Conn, theMessage string) {
	botConnection := botConn

	theCommand := strings.Split(strings.TrimSpace(theMessage), " ")

	Command := strings.TrimPrefix(theCommand[0], ">")

	handleMessage(Command, theCommand, botConnection)
	// if theMessage == "hello" {
	// 	_, err := botConnection.Write([]byte("hello"))
	// 	if err != nil {
	// 		Misc.OutputLogger.Fatalln("Error When sending", err)
	// 	}
	// }
}

func handleMessage(command string, theCommand []string, c net.Conn) {
	switch command {
	case "setbotname":
		theBot := data.Bot{
			FirstName:  theCommand[1],
			Connection: c,
		}
		err := theBot.AddBot()
		if err != nil {
			misc.Info("Can't Add Bot >> ", err)
		}
	default:
		//Printout to the output logger
		fmt.Println(strings.Join(theCommand,""))
	}
}