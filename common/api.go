package common

import (
	"fmt"
	"net/http"
)

func init() {
	// api.go init
}

// parser table stats
func ParserTs(StatusPort int, DbIp string, dbName string, tbName string) {
	req, err := http.Get(fmt.Sprintf("http://%s:%v/stats/dump/%s/%s", DbIp, StatusPort, dbName, tbName))
	if err != nil {
		//
	}
	defer req.Body.Close()
	fmt.Println(req)
}
