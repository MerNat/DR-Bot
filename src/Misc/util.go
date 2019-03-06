package misc

import (
	"log"
	"os"
)

var fileLogger *log.Logger
var outputLogger *log.Logger

func init() {
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Can not open logger file", err)
	}

	outputLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	fileLogger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfi(){
	
}

// Danger logs only danger infos to loggers.
func Danger(msg...interface{}) {
	outputLogger.SetPrefix("DANGER ")
	outputLogger.Println(msg...)
	fileLogger.SetPrefix("DANGER ")
	fileLogger.Println(msg...)
}

// Warning logs only warning infos to loggers.
func Warning(msg...interface{}) {
	outputLogger.SetPrefix("WARNING ")
	outputLogger.Println(msg...)
	fileLogger.SetPrefix("WARNING ")
	fileLogger.Println(msg...)
}

// Err function prints the err and exit with err code 1
func Err(msg...interface{}) {
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
