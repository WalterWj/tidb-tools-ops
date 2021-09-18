package common

import "log"

// error log
func IfErrLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
