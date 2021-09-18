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
func IfErrPrint(errInfo string) {
	t := time.Unix(0, time.Now().UnixMilli()*1000000)
	ct := t.Format(time.RFC3339Nano)
	fmt.Printf("%s ERROR: %s\n", ct, errInfo)
}
