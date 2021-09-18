package common

import (
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
	tableList := strings.Split(tbargs, ",")
	tableMap := make(map[string][]string)
	var tbTmp []string
	for _, tb := range tableList {
		part := strings.Split(tb, ".")
		dbName := part[0]
		tbTmp = append(tableMap[dbName], part[1])
		tableMap[dbName] = tbTmp
	}
	return tableMap
}
