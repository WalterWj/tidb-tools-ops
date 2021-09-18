package common

import (
	"fmt"
	"log"
	"time"
)

// error log
func IfErrLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// error print
func IfErrPrint(err string) {
	ct := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("%s ERROR: %s\n", ct, err)
}
