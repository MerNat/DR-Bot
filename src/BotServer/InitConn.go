package main

import (
	botcommand "BotCommand"
	misc "Misc"
	"bufio"
	"data"
	"fmt"
	"net"
	"os"
	"strings"
)

// InitConn initialize the connection
func InitConn(netType string, address string, port string) {
	myConn, err := net.Listen(netType, address+":"+port)
	if err != nil {
		misc.Err("Error Creating Server!", err)
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
			bot, err := data.GetBotByConnection(c)
			if err == nil{
				bot.Remove()
				misc.Info(bot.FirstName," Bot has left room")
				return
			}else{
			}
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

// InitInput handles the input from the server
func InitInput() {
	misc.Info("Input from server has initialized")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			misc.Err("Can not read from input", err)
		}
		handleInput(text)
	}
}

func handleInput(msgFromServer string) {
	theCommand := strings.Split(strings.TrimSpace(msgFromServer), " ")
	forAll := strings.IndexAny(theCommand[0], ">")
	if forAll == -1 {
		//Send to all
		handleInputCommands("", theCommand)
	} else {
		//Send to specific Bot
		botName := strings.TrimPrefix(theCommand[0], ">")
		handleInputCommands(botName, theCommand)
	}
}

func handleInputCommands(botName string, theCommand []string) {
	if botName != "" {
		//Send it to specific Bot
		bot, err := data.GetBotByName(botName)
		if err != nil {
			misc.Warning(err)
		} else {
			bot.SendBotMsg(strings.Join(theCommand, ""))
		}
	}
}
