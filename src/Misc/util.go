package misc

import (
	"encoding/json"
	"log"
	"os"
)

var (
	fileLogger   *log.Logger
	outputLogger *log.Logger
	clientConfig       ClientConfiguration
)

//ClientConfiguration the client config details
type ClientConfiguration struct {
	BOT_URL      string
	CHAT_ID   string
	CHAT_TYPE string
}

//ServerConfig struct
type ServerConfig struct {
	Address string
	Port    string
	NetType string
}

//ServerConfiguration has the required configuration for the server
var ServerConfiguration ServerConfig

func init() {
	loadConfig()
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Can not open logger file", err)
	}

	outputLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	fileLogger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.OpenFile("config.json", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln("Can't open Configuration file.", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ServerConfiguration)
	if err!=nil{
		log.Fatalln("Can't decode config file.", err)
	}
}

// Danger logs only danger infos to loggers.
func Danger(msg ...interface{}) {
	outputLogger.SetPrefix("DANGER ")
	outputLogger.Println(msg...)
	fileLogger.SetPrefix("DANGER ")
	fileLogger.Println(msg...)
}

// Warning logs only warning infos to loggers.
func Warning(msg ...interface{}) {
	outputLogger.SetPrefix("WARNING ")
	outputLogger.Println(msg...)
	fileLogger.SetPrefix("WARNING ")
	fileLogger.Println(msg...)
}

// Err function prints the err and exit with err code 1
func Err(msg ...interface{}) {
	fileLogger.SetPrefix("ERROR ")
	fileLogger.Println(msg...)
	outputLogger.SetPrefix("ERROR ")
	outputLogger.Fatalln(msg...)
}

// Info logs information.
func Info(msg ...interface{}) {
	outputLogger.SetPrefix("INFO ")
	outputLogger.Println(msg...)
	fileLogger.SetPrefix("INFO ")
	fileLogger.Println(msg...)
}
