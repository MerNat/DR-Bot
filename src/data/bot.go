package data

import (
	misc "Misc"
	"errors"
	"net"
	"sort"
)

// Bot structure implements what a Bot should have as a Client Bot.
type Bot struct {
	// id            int64  `json:"id"`
	FirstName  string `json:"firstName"`
	Connection net.Conn
}

// Bots has every Bot inside it
var Bots []Bot

// Message is an object that will be sent to the Bot.

// Chat is an object that tells the DRBot it's for a channel
// type Chat struct{
// 	id	int64 `json:"id"`
// 	chatType string `json:"ChatType"`
// }
// type Message struct{

// }

// AddBot adds the Bot
func (bot *Bot) AddBot() (err error) {
	if findBot(bot.FirstName) {
		//Should notify the Bot to generate a new name
		bot.SendBotMsg("Name Taken")
		err = errors.New("Name Taken Generated")
	} else {
		//Add the Bot
		Bots = append(Bots, *bot)
		misc.Info(bot.FirstName, " Bot Joined")
	}
	return
}

func sortSlice() {
	sort.SliceStable(Bots, func(i, j int) bool { return Bots[i].FirstName < Bots[j].FirstName })
}

func findBot(firstName string) (isThere bool) {
	isThere = false
	sortSlice()
	for _, bot := range Bots {
		if bot.FirstName == firstName {
			isThere = true
			break
		}
	}
	return
}

// SendBotMsg sends a message to a bot
func (bot *Bot) SendBotMsg(msg string) {
	_, err := bot.Connection.Write([]byte(msg))
	if err != nil {
		misc.Danger("Error When sending", err)
	}
}

// GetBotByName returns the bot by name
func GetBotByName(botName string) (theBot Bot, err error) {
	isBot := false
	for _, bot := range Bots {
		if bot.FirstName == botName {
			theBot = bot
			isBot = true
			break
		}
	}
	if !isBot {
		err = errors.New("No bot with such name")
	}
	return
}

//GetBotByConnection returns a Bot struct based on a net.Conn
func GetBotByConnection(conn net.Conn) (theBot Bot, err error) {
	isBot := false
	for _, bot := range Bots {
		if bot.Connection == conn {
			theBot = bot
			isBot = true
			break
		}
	}
	if !isBot {
		err = errors.New("No bot with such connection")
	}
	return
}

//RemoveBot remove and deletes the bot from the slice Bots
func (bot *Bot) RemoveBot() {

}

//GetBotsIndex returns the bots index in Bots slice
func (bot *Bot) GetBotsIndex() {

}
