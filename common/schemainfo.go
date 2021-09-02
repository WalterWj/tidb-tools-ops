package common

import (
	"database/sql"
	"fmt"
)

func init() {
	// fmt.Println("get schema information mould init funcation")
}

// Get table schema information
func GetTables(db *sql.DB, dbname string) map[int]string {
	var r = make(map[int]string)
	tablesQ := fmt.Sprintf("select table_name from information_schema.tables where TABLE_SCHEMA in (%s);", dbname)
	rows, err := db.Query(tablesQ)
	if err != nil {
		fmt.Printf("execute %v fail", tablesQ)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			fmt.Printf("rows scan fail")
		}
		r[n] = t
		n++
	}
	return r
}
