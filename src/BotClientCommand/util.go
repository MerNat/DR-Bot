package botclientcommand

import (
	m "MiscClient"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

// SendToServer sends a message to server
func sendToServer(c net.Conn, msg string) (err error) {
	_, err = c.Write([]byte(msg))
	return
}

func sendMessageToChannel(message string) {
	message = strings.Replace(message, " ", "", -1)
	resp, err := http.Get(m.ClientConfig.BOT_URL + "sendMessage?chat_id=" + m.ClientConfig.CHAT_ID + "&text=@" + BotName+":" + message)
	if err != nil {
		log.Fatalln("cannot read url body")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}
