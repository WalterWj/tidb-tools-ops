package argsutil

import (
	"fmt"
	"strings"
	log "tidb-tools-ops/pkg/logutil"
)

func init() {
	// fmt.Println("file mould init funcation")
}

// parser db args: test, test1, test2
func ParserDbArgs(dbargs string) []string {
	dbList := strings.Split(dbargs, ",")
	return dbList
}

// parser table args: test1.t1, test1.t2, test2,t1
func ParserTbArgs(tbargs string) map[string][]string {
	tableMap := make(map[string][]string)
	var tbTmp []string
	tableList := strings.Split(tbargs, ",")
	for _, tb := range tableList {
		// parser table args
		part := strings.Split(tb, ".")
		dbName := part[0]
		// 判断 table args 是否有问题
		if len(part) < 2 {
			log.ErrorLog(fmt.Sprintf("table args: %s is wrong", tb))
		} else {
			// add table name
			tbTmp = append(tableMap[dbName], part[1])
			// make map[dbname]tbname
			tableMap[dbName] = tbTmp
		}
	}

	return tableMap
}

// map[string]string to []string
func MapToArryString(maplist map[string]string) []string {
	var arryString []string
	for _, _string := range maplist {
		arryString = append(arryString, _string)
	}
	return arryString
}
