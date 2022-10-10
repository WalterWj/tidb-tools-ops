package dbutil

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Tables struct {
	Dbname     string
	Table_name string
}

func TableList(db *sql.DB, dbname string) (tables *Tables, err error) {
	tables = &Tables{}
	tables.Dbname = dbname
	// Determine whether the database exists
	ok := IfDbNotE(db, tables.Dbname)
	if ok {
		// get tables name
		tablesQ := fmt.Sprintf("select table_name from information_schema.tables where TABLE_SCHEMA in (%s) and TABLE_TYPE <> 'VIEW' order by 1;", strconv.Quote(tables.Dbname))
		rows, _ := db.Query(tablesQ)
		//
		for rows.Next() {
			rows.Scan(&tables.Table_name)
			fmt.Printf("tables: %+v\n", tables)
		}

		defer func() {
			rows.Close()
		}()

	} else {
		fmt.Printf("[WARN] Database %s is not exist  \n", dbname)
	}

	return tables, nil
}
