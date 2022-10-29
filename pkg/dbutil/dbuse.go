package dbutil

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	log "tidb-tools-ops/pkg/logutil"
)

// Query return []map[string]string, error
func Query(db *sql.DB, SQL string) ([]map[string]string, error) {
	// make map,  创建返回值：不定长的map类型切片
	ret := make([]map[string]string, 0)
	// execute ques
	rows, err := db.Query(SQL)
	if err != nil {
		log.ErrorLog(err.Error())
		return ret, errors.New("Query fail,\n" + err.Error())
	}
	// get columns info
	columns, _ := rows.Columns()
	// columns lenth
	count := len(columns)
	// values
	var values = make([]interface{}, count)
	// 为空接口分配内存
	for i := range values {
		var v interface{}
		values[i] = &v
	}
	for rows.Next() {
		//开始读行，Scan函数只接受指针变量
		err := rows.Scan(values...)
		//用于存放1列的 [键/值] 对
		m := make(map[string]string)
		if err != nil {
			log.ErrorLog(err.Error())
			return ret, errors.New("row scan fail\n" + err.Error())
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
		return ret, nil
	}

	return ret, errors.New("effect 0 rows")
}

// if database is not exist
func IfDbNotE(db *sql.DB, dbname string) bool {
	dbQ := fmt.Sprintf("select SCHEMA_NAME as c from information_schema.SCHEMATA where SCHEMA_NAME in (%s);", strconv.Quote(dbname))
	_, OK := Query(db, dbQ)
	// 如果库存在，返回为 true，OK 为 true； 如果不存在，OK 为 false，返回为 false
	if OK == nil {
		return true
	} else {
		return false
	}
}

// Get table name
func GetTables(db *sql.DB, dbname string) map[string]string {
	var r = make(map[string]string)
	// Determine whether the database exists
	ok := IfDbNotE(db, dbname)
	if ok {
		// get tables name
		tablesQ := fmt.Sprintf("select table_name from information_schema.tables where TABLE_SCHEMA in (%s) and TABLE_TYPE <> 'VIEW';", strconv.Quote(dbname))
		sr, ok := Query(db, tablesQ)
		if ok == nil {
			for _, _sr := range sr {
				r[_sr["table_name"]] = _sr["table_name"]
			}
			return r
		} else {
			fmt.Printf("execute %v fail\n", tablesQ)
		}
	} else {
		// fmt.Printf("[WARN] Database %s is not exist  \n", dbname)
		log.WarningLog(fmt.Sprintf("[WARN] Database %s is not exist  \n", dbname))
	}

	return r
}

// get states healthy for table
func GetTableHealthy(db *sql.DB, dbname string, tablename string, healthy int) bool {
	var r bool
	// get healthy
	_gh_sql := fmt.Sprintf("show stats_healthy where Db_name in (%s) and Table_name in (%s);", strconv.Quote(dbname), strconv.Quote(tablename))
	rows, ok := Query(db, _gh_sql)
	if ok == nil {
		for _, _rc := range rows {
			_result := _rc["Healthy"]
			health, _ := strconv.Atoi(_result)
			if health <= healthy {
				r = true
			} else {
				log.WarningLog(fmt.Sprintf("DB: %s,Table: %s,healthy: %d", dbname, tablename, health))
				r = false
			}
		}
		return r
	} else {
		fmt.Printf("execute %v fail\n", _gh_sql)
	}
	return r
}

// get db sql for mode
func GetDbSql(mode int) string {
	if mode == 0 {
		tablesQ := `select distinct TABLE_SCHEMA from INFORMATION_SCHEMA.tables where TABLE_SCHEMA 
		not in ('METRICS_SCHEMA','PERFORMANCE_SCHEMA','INFORMATION_SCHEMA','mysql');`
		return tablesQ
	} else if mode == 1 {
		tablesQ := `select distinct TABLE_SCHEMA from INFORMATION_SCHEMA.tables;`
		return tablesQ
	} else {
		err := "Please input 0/1 for mode"
		log.ErrorLog(err)
		return err
	}
}

// get db name,ignore 'METRICS_SCHEMA','PERFORMANCE_SCHEMA','INFORMATION_SCHEMA','mysql'
func GetAllDb(db *sql.DB, mode int) map[string]string {
	var r = make(map[string]string)
	tablesQ := GetDbSql(mode)
	rows, ok := Query(db, tablesQ)
	if ok == nil {
		for _, _rc := range rows {
			r[_rc["TABLE_SCHEMA"]] = _rc["TABLE_SCHEMA"]
		}
		return r
	} else {
		// fmt.Printf("execute %v fail\n", tablesQ)
		log.ErrorLog(fmt.Sprintf("execute %v fail\n", tablesQ))
	}

	return r
}

// get TiDB version
func GetVersion(db *sql.DB) map[int]string {
	var r = make(map[int]string)
	const Query = "select tidb_version();"
	rows, err := db.Query(Query)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("execute %v fail\n", Query))
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		if err != nil {
			log.ErrorLog("GetVersion, rows scan fail\n")
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
		log.ErrorLog(fmt.Sprintf("execute %v fail\n", tablesQ) + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var t, ct string
		err := rows.Scan(&t, &ct)
		if err != nil {
			log.ErrorLog(tablesQ + "ParserTables, rows scan fail\n")
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
		log.ErrorLog(fmt.Sprintf("execute %v fail\n", DbQ))
	}
	defer rows.Close()
	for rows.Next() {
		var d, cd string
		err := rows.Scan(&d, &cd)
		if err != nil {
			log.ErrorLog(DbQ + "ParserDb, rows scan fail\n")
		}
		r[d] = cd
	}
	return r[dbname]
}
