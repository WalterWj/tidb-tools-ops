package common

import (
	"fmt"
	"strings"
)

func init() {
	// fmt.Println("file mould init funcation")
}

// parser db args
func ParserDbArgs(dbargs string) []string {
	dbList := strings.Split(dbargs, ",")
	return dbList
}

// parser table args
func ParserTbArgs(tbargs string) map[string][]string {
	tableMap := make(map[string][]string)
	var tbTmp []string
	tableList := strings.Split(tbargs, ",")
	for _, tb := range tableList {
		// parser table args
		part := strings.Split(tb, ".")
		dbName := part[0]

		if len(part) < 2 {
			err := fmt.Sprintf("table args: %s is wrong", tb)
			IfErrPrintE(err)
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
