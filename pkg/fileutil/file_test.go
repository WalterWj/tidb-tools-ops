package fileutil_test

import (
	"os"
	"testing"
	"tidb-tools-ops/pkg/logutil"
)

func TestAddfil(t *testing.T) {
	logutil.InitLog("test.log")
	defer os.Remove("test.log")
	// err := logutil.InitLog("/root/test.test")
	// want := "open /root/test.test: permission denied"
	// if want != err.Error() {
	// 	t.Fatalf(`InitLog("/root/test.test") = %s, want "", error`, err.Error())
	// }
}
