package common

import (
	"fmt"
	"log"
	"strconv"
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
	i, err := strconv.ParseInt("1489582166", 10, 64)
	if err != nil {
		panic(err)
	}
	ct := time.Unix(i, 0)
	fmt.Printf("%s ERROR: %s\n", ct, errInfo)
}
