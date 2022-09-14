package logutil

import (
	"fmt"
	"log"
	"os"
)

// var warning/info/error log
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// init log file
func InitLog(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}

	InfoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "[WARNING] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime)

	return nil
}

// info log
func InfoLog(c string) {
	InfoLogger.Println(c)
}

// error log
func ErrorLog(c string) {
	ErrorLogger.Println(c)
}

// warning log
func WarningLog(c string) {
	WarningLogger.Println(c)
}
