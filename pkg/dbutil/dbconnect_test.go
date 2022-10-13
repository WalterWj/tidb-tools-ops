package dbutil_test

import (
	"fmt"
	"strings"
	"testing"
	"tidb-tools-ops/pkg/dbutil"
	log "tidb-tools-ops/pkg/logutil"
)

func TestConnectDB(t *testing.T) {
	log.InitLog("test.log")
	// nomal
	dsn := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(4201), ")/", "mysql", "?charset=utf8"}, "")
	_, err := dbutil.MysqlConnect(dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("connect sucessful")

	// db port wrong
	dsn1 := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(42011), ")/", "mysql", "?charset=utf8"}, "")
	_, err = dbutil.MysqlConnect(dsn1)
	if err != nil {
		fmt.Println(err.Error())
	}
	want := "Ping MySQL fail~\ndial tcp 127.0.0.1:42011: connect: connection refused"
	if want != err.Error() {
		t.Fatalf("errors: %s", err.Error())
	}
}
