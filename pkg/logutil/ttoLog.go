package logutil

import (
	"log"
	"os"
)

// interface
type Logs interface {
	// init log file
	InitLog()
	// write info log file
	InfoLog()
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type Content struct {
	Content string
	Mode    string `default:"info"`
}

func InitLog(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InfoLog(c string) {
	InfoLogger.Println(c)
}

func ErrorLog(c string) {
	ErrorLogger.Println(c)
}

func WarningLog(c string) {
	WarningLogger.Println(c)
}

// use
// func InfoLog() {
// 	InfoLogger.Println("Starting the application...")
// }
