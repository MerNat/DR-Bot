package Misc

import (
	"log"
	"os"
)

var FileLogger *log.Logger
var OutputLogger *log.Logger

func init() {
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalln("Can not open logger file", err)
	}

	defer file.Close()

	OutputLogger = log.New(os.Stdout,"INFO ",log.Ldate | log.Ltime | log.Lshortfile)
	FileLogger = log.New(file, "INFO ",log.Ldate | log.Ltime | log.Lshortfile)
}

func danger(msg string) {
	OutputLogger.SetPrefix("DANGER ")
	OutputLogger.Println(msg)
	FileLogger.SetPrefix("DANGER ")
	FileLogger.Println(msg)
}

func warning(msg string) {
	OutputLogger.SetPrefix("WARNING ")
	OutputLogger.Println(msg)
	FileLogger.SetPrefix("WARNING ")
	FileLogger.Println(msg)
}

func err(msg string, err error) {
	FileLogger.SetPrefix("ERROR ")
	FileLogger.Println(msg, err.Error())
	OutputLogger.SetPrefix("ERROR ")
	OutputLogger.Fatalln(msg, err.Error())
	
}

func info(msg string) {
	OutputLogger.SetPrefix("INFO ")
	OutputLogger.Println(msg)
	FileLogger.SetPrefix("INFO ")
	FileLogger.Println(msg)
}
