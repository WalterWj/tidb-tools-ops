package common

import (
	"database/sql"
	"fmt"
	"strconv"
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
	err = db.Ping()
	if err != nil {
		IfErrPrintE("Connect MySQL fail~")
	}
	return db
}

// Query
func Query(db *sql.DB, SQL string) ([]map[string]string, bool) {
	// execute ques
	rows, err := db.Query(SQL)
	if err != nil {
		IfErrLog(err)
	}
	// get columns info
	columns, _ := rows.Columns()
	// columns lenth
	count := len(columns)
	// values
	var values = make([]interface{}, count)
	// 为空接口分配内存
	for i, _ := range values {
		var v interface{}
		values[i] = &v
	}
	// make map,  创建返回值：不定长的map类型切片
	ret := make([]map[string]string, 0)
	for rows.Next() {
		//开始读行，Scan函数只接受指针变量
		err := rows.Scan(values...)
		//用于存放1列的 [键/值] 对
		m := make(map[string]string)
		if err != nil {
			IfErrLog(err)
		}
		for i, colName := range columns {
			// 读出raw数据，类型为byte
			var raw_value = *(values[i].(*interface{}))
			b, _ := raw_value.([]byte)
			//将raw数据转换成字符串
			v := string(b)
			// colName是键，v是值
			m[colName] = v
		}
		// 将单行所有列的键值对附加在总的返回值上（以行为单位）
		ret = append(ret, m)
	}

	defer rows.Close()

	if len(ret) != 0 {
		return ret, true
	}

	return nil, false
}

// if database is not exist
func IfDbNotE(db *sql.DB, dbname string) int {
	dbQ := fmt.Sprintf("select SCHEMA_NAME as c from information_schema.SCHEMATA where SCHEMA_NAME in (%s);", strconv.Quote(dbname))
	_, OK := Query(db, dbQ)
	// 如果库存在，返回为 1，OK 为 true； 如果不存在，OK 为 false，返回为 0
	if OK {
		return 1
	} else {
		return 0
	}
}

// Get table name
func GetTables(db *sql.DB, dbname string) map[int]string {
	var r = make(map[int]string)
	// Determine whether the database exists
	rc := IfDbNotE(db, dbname)
	if rc == 0 {
		fmt.Printf("[WARN] Database %s is not exist  \n", dbname)
	}
	// get tables name
	tablesQ := fmt.Sprintf("select table_name from information_schema.tables where TABLE_SCHEMA in (%s) and TABLE_TYPE <> 'VIEW';", strconv.Quote(dbname))
	rows, err := db.Query(tablesQ)
	if err != nil {
		fmt.Printf("execute %v fail\n", tablesQ)
	}
	defer rows.Close()
	// make tables: map[int]string, {1: test, 2: test1}
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			fmt.Printf("rows scan fail\n")
			IfErrPrintE(tablesQ)
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
		IfErrLog(err)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			IfErrPrintE("rows scan fail\n")
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
		IfErrLog(err)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			IfErrPrintE("GetVersion, rows scan fail\n")
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
		IfErrLog(err)
	}
	defer rows.Close()
	for rows.Next() {
		var t, ct string
		err := rows.Scan(&t, &ct)
		if err != nil {
			fmt.Printf("ParserTables, rows scan fail\n")
			IfErrPrintE(tablesQ)
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
			IfErrPrintE(DbQ)
		}
		r[d] = cd
	}
	return r[dbname]
}
