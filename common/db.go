package common

import (
	"database/sql"
	"fmt"
)

func init() {
	// fmt.Println("get schema information mould init funcation")
}

// Get table name
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

// get db sql
func getDbSql(mode string) string {
	if len(mode) == 0 {
		tablesQ := `select distinct TABLE_SCHEMA from tables where TABLE_SCHEMA 
		not in ('METRICS_SCHEMA','PERFORMANCE_SCHEMA','INFORMATION_SCHEMA','mysql');`
		return tablesQ
	} else {
		tablesQ := `select distinct TABLE_SCHEMA from tables`
		return tablesQ
	}
}

// get db name,ignore 'METRICS_SCHEMA','PERFORMANCE_SCHEMA','INFORMATION_SCHEMA','mysql'
func getAllDb(db *sql.DB, dbname string, mode string) map[int]string {
	var r = make(map[int]string)
	tablesQ := getDbSql(mode)
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

// get TiDB version
func GetVersion(db *sql.DB) map[int]string {
	var r = make(map[int]string)
	const Query = "select tidb_version();"
	rows, err := db.Query(Query)
	if err != nil {
		fmt.Printf("execute %v fail", Query)
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

// get table schema
func ParserTables(db *sql.DB, dbname string, tablename string) string {
	var r = make(map[string]string)
	tablesQ := fmt.Sprintf("show create table `%v`.`%v`;", dbname, tablename)
	rows, err := db.Query(tablesQ)
	if err != nil {
		fmt.Printf("execute %v fail", tablesQ)
	}
	defer rows.Close()
	for rows.Next() {
		var t, ct string
		err := rows.Scan(&t, &ct)
		if err != nil {
			fmt.Printf("rows scan fail")
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
		fmt.Printf("execute %v fail", DbQ)
	}
	defer rows.Close()
	for rows.Next() {
		var d, cd string
		err := rows.Scan(&d, &cd)
		if err != nil {
			fmt.Printf("rows scan fail")
		}
		r[d] = cd
	}
	return r[dbname]
}
