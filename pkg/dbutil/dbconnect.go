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
