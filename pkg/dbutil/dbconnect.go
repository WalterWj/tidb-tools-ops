package dbutil

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	log "tidb-tools-ops/pkg/logutil"
)

// db connect
func MysqlConnect(dsn string) (*sql.DB, error) {
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.ErrorLog("connect db fail." + err.Error())
		return nil, errors.New("connect db fail.\n" + err.Error())
	}

	err = Db.Ping()
	if err != nil {
		log.ErrorLog("Ping MySQL fail~")
		return nil, errors.New("Ping MySQL fail~\n" + err.Error())
	}

	return Db, nil
}

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
