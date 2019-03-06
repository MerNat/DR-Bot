package botcommand

import (
	"fmt"
	"net"
	"strings"
)

var botConnection net.Conn

// ReadBotMessage Reads the message comming from a bot and make it ready to process
func ReadBotMessage(botConn net.Conn, theMessage string) {
	fmt.Print(theMessage)
	botConnection = botConn
	theCommand := strings.Split(strings.TrimSpace(theMessage), " ")

	Command := strings.TrimPrefix(theCommand[0], ">")

	handleMessage(Command, theCommand)
	// if theMessage == "hello" {
	// 	_, err := botConnection.Write([]byte("hello"))
	// 	if err != nil {
	// 		Misc.OutputLogger.Fatalln("Error When sending", err)
	// 	}
	// }
}

func handleMessage(command string, theCommand []string){
	switch command{
	case "setbotname":
		// botName := theCommand[1]
	}
}

func sendMessage(c net.Conn){

}