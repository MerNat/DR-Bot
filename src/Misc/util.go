package misc

import (
	"log"
	"os"
)

var FileLogger *log.Logger
var OutputLogger *log.Logger

func init() {
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Can not open logger file", err)
	}

	OutputLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	FileLogger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func danger(msg...interface{}) {
	OutputLogger.SetPrefix("DANGER ")
	OutputLogger.Println(msg...)
	FileLogger.SetPrefix("DANGER ")
	FileLogger.Println(msg...)
}

func warning(msg...interface{}) {
	OutputLogger.SetPrefix("WARNING ")
	OutputLogger.Println(msg...)
	FileLogger.SetPrefix("WARNING ")
	FileLogger.Println(msg...)
}

func err(msg...interface{}) {
	FileLogger.SetPrefix("ERROR ")
	FileLogger.Println(msg...)
	OutputLogger.SetPrefix("ERROR ")
	OutputLogger.Fatalln(msg...)

}

// Info logs information.
func Info(msg ...interface{}) {
	OutputLogger.SetPrefix("INFO ")
	OutputLogger.Println(msg...)
	FileLogger.SetPrefix("INFO ")
	FileLogger.Println(msg...)
}
