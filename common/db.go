package common

import (
	"database/sql"
	"fmt"
)

func init() {
	// fmt.Println("get schema information mould init funcation")
}

// db connect
func MysqlConnect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db fail. %s", err)
	}
	return db
}

// Get table name
func GetTables(db *sql.DB, dbname string) map[int]string {
	var r = make(map[int]string)
	tablesQ := fmt.Sprintf("select table_name from information_schema.tables where TABLE_SCHEMA in (%s) and TABLE_TYPE <> 'VIEW';;", dbname)
	rows, err := db.Query(tablesQ)
	if err != nil {
		fmt.Printf("execute %v fail\n", tablesQ)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			fmt.Printf("rows scan fail\n")
			IfErrPrint(tablesQ)
		}
		r[n] = t
		n++
	}
	return r
}

// get db sql
func GetDbSql(mode int) string {
	if mode == 0 {
		tablesQ := `select distinct TABLE_SCHEMA from INFORMATION_SCHEMA.tables where TABLE_SCHEMA 
		not in ('METRICS_SCHEMA','PERFORMANCE_SCHEMA','INFORMATION_SCHEMA','mysql');`
		return tablesQ
	} else if mode == 1 {
		tablesQ := `select distinct TABLE_SCHEMA from INFORMATION_SCHEMA.tables;`
		return tablesQ
	} else {
		panic("Please input 0/1 for mode")
	}
}

// get db name,ignore 'METRICS_SCHEMA','PERFORMANCE_SCHEMA','INFORMATION_SCHEMA','mysql'
func GetAllDb(db *sql.DB, mode int) map[int]string {
	var r = make(map[int]string)
	tablesQ := GetDbSql(mode)
	rows, err := db.Query(tablesQ)
	if err != nil {
		fmt.Printf("execute %v fail\n", tablesQ)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			IfErrPrint("rows scan fail\n")
		}
		r[n] = t
		n++
	}
	return r
}

// get TiDB version
func GetVersion(db *sql.DB) map[int]string {
	var r = make(map[int]string)
	const Query = "select tidb_version();"
	rows, err := db.Query(Query)
	if err != nil {
		fmt.Printf("execute %v fail\n", Query)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			IfErrPrint("GetVersion, rows scan fail\n")
		}
		r[n] = t
		n++
	}
	return r
}

// get table schema
func ParserTables(db *sql.DB, dbname string, tablename string) string {
	var r = make(map[string]string)
	tablesQ := fmt.Sprintf("show create table `%v`.`%v`;", dbname, tablename)
	rows, err := db.Query(tablesQ)
	if err != nil {
		fmt.Printf("execute %v fail\n", tablesQ)
	}
	defer rows.Close()
	for rows.Next() {
		var t, ct string
		err := rows.Scan(&t, &ct)
		if err != nil {
			fmt.Printf("ParserTables, rows scan fail\n")
			IfErrPrint(tablesQ)
		}
		r[t] = ct
	}
	return r[tablename]
}

// get database schema
func ParserDb(db *sql.DB, dbname string) string {
	var r = make(map[string]string)
	DbQ := fmt.Sprintf("show create database if not exists `%v`;", dbname)
	rows, err := db.Query(DbQ)
	if err != nil {
		fmt.Printf("execute %v fail\n", DbQ)
	}
	defer rows.Close()
	for rows.Next() {
		var d, cd string
		err := rows.Scan(&d, &cd)
		if err != nil {
			fmt.Printf("ParserDb, rows scan fail\n")
			IfErrPrint(DbQ)
		}
		r[d] = cd
	}
	return r[dbname]
}
