package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	// api.go init
}

// parser table stats
func ParserTs(DbIp string, StatusPort int, dbName string, tbName string) string {
	Apipath := fmt.Sprintf("http://%s:%v/stats/dump/%s/%s", DbIp, StatusPort, dbName, tbName)
	response, err := http.Get(Apipath)
	if err != nil {
		fmt.Printf("curl %s faild", Apipath)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf(err.Error())
	}
	content := string(body)
	return content
}
