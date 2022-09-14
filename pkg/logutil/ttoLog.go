package logutil

import (
	"fmt"
	"log"
	"os"
)

// interface
type Logs interface {
	// init log file
	InitLog()
	// write info log file
	InfoLog()
	// write error log file
	ErrorLog()
	// write warning log file
	WarningLog()
}

// var warning/info/error log
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// log content
type Content struct {
	Content string
}

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
func (c Content) InfoLog() {
	InfoLogger.Println(c.Content)
}

// error log
func (c Content) ErrorLog() {
	ErrorLogger.Println(c.Content)
}

// warning log
func (c Content) WarningLog() {
	WarningLogger.Println(c.Content)
}
