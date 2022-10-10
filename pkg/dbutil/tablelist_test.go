package dbutil_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"tidb-tools-ops/pkg/dbutil"
	"tidb-tools-ops/pkg/logutil"
)

func TestAddfil(t *testing.T) {
	logutil.InitLog("test.log")
	defer os.Remove("test.log")
	dsn := strings.Join([]string{"root", ":", "tidb@123", "@tcp(", "127.0.0.1", ":", fmt.Sprint(4201), ")/", "mysql", "?charset=utf8"}, "")
	db, _ := dbutil.MysqlConnect(dsn)
	tmp, _ := dbutil.TableList(db, "test")
	// for table_name, Id := range tmp {
	// 	fmt.Println(table_name, Id)
	// }
	// for i range tmp{}
	fmt.Println(tmp)
	// err := logutil.InitLog("/root/test.test")
	// want := "open /root/test.test: permission denied"
	// if want != err.Error() {
	// 	t.Fatalf(`InitLog("/root/test.test") = %s, want "", error`, err.Error())
	// }
}
