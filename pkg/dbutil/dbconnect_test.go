package dbutil

import (
	"fmt"
	"strings"
	"testing"
	log "tidb-tools-ops/pkg/logutil"
)

func TestConnectDB(t *testing.T) {
	log.InitLog("test.log")
	// nomal
	dsn := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(4201), ")/", "mysql", "?charset=utf8"}, "")
	_, err := MysqlConnect(dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("connect sucessful")

	// db port wrong
	dsn1 := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(42011), ")/", "mysql", "?charset=utf8"}, "")
	_, err = MysqlConnect(dsn1)
	if err != nil {
		fmt.Println(err.Error())
	}
	want := "Ping MySQL fail~\ndial tcp 127.0.0.1:42011: connect: connection refused"
	if want != err.Error() {
		t.Fatalf("errors: %s", err.Error())
	}
}

func TestDBQuery(t *testing.T) {
	log.InitLog("test.log")

	dsn := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(4201), ")/", "mysql", "?charset=utf8"}, "")
	db, err := MysqlConnect(dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	// test nomal
	rst, err := Query(db, "select 1;")
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
	rst1, err := Query(db, "select a;")
	if err != nil {
		fmt.Println("erros:", err.Error())
	}
	var want1 = string("Query fail,\nError 1054: Unknown column 'a' in 'field list'")
	if want1 != err.Error() {
		t.Fatal("err: ", err.Error())
	}
	fmt.Println("result:", rst1)

	// test 0 rows
	rst2, err := Query(db, "select * from test.tmp;")
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
