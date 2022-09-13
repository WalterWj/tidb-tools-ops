package dbutil

import (
	"database/sql"
	"fmt"
)

// db connect
func MysqlConnect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db fail. %s", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Connect MySQL fail~")
	}
	return db
}
