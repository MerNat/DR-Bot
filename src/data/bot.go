package data

// Bot structure implements what a Bot should have as a Client Bot.
type Bot struct {
	id         int64  `json:"id"`
	firstName string `json:"firstName"`
}

// Message is an object that will be sent to the Bot.


// Chat is an object that tells the DRBot it's for a channel
// type Chat struct{
// 	id	int64 `json:"id"`
// 	chatType string `json:"ChatType"`
// }
// type Message struct{

// }