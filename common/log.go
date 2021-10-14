package common

import (
	"fmt"
	"log"
	"time"
)

// error log, Exit when an error is reported
func IfErrLog(err error) {
	if err != nil {
		fmt.Print("[ERROR] ")
		log.Fatal(err)
	}
}

// error print, Continue when an error is reported
func IfErrPrintE(errInfo string) {
	t := time.Unix(0, time.Now().UnixMilli()*1000000)
	ct := t.Format(time.RFC3339Nano)
	fmt.Printf("[%s] [Error] %s\n", ct, errInfo)
}
