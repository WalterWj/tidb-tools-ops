package dbutil_test

import (
	"fmt"
	"strings"
	"testing"
	"tidb-tools-ops/pkg/dbutil"
	log "tidb-tools-ops/pkg/logutil"
)

func TestDBQuery(t *testing.T) {
	log.InitLog("test.log")

	dsn := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(4201), ")/", "mysql", "?charset=utf8"}, "")
	db, err := dbutil.MysqlConnect(dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	// test nomal
	rst, err := dbutil.Query(db, "select 1;")
	if err != nil {
		fmt.Println(err.Error())
	}

	// var want []map[string]string
	var want = map[string]string{
		"1": "1",
	}
	if want["1"] != rst[0]["1"] {
		t.Fatalf("want: %s, result: %s", want, rst)
	}
	fmt.Println("result:", rst[0]["1"], "want:", want["1"])

	// test error
	rst1, err := dbutil.Query(db, "select a;")
	if err != nil {
		fmt.Println("erros:", err.Error())
	}
	var want1 = string("Query fail,\nError 1054: Unknown column 'a' in 'field list'")
	if want1 != err.Error() {
		t.Fatal("err: ", err.Error())
	}
	fmt.Println("result:", rst1)

	// test 0 rows
	rst2, err := dbutil.Query(db, "select * from test.tmp;")
	if err != nil {
		fmt.Println("erros:", err.Error())
	}
	var want2 = string("effect 0 rows")
	if want2 != err.Error() {
		t.Fatal("err: ", err.Error())
	}
	fmt.Println("result:", rst2)

	defer db.Close()
}
