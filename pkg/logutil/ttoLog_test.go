package logutil_test

import (
	"fmt"
	"os"
	"testing"
	"tidb-tools-ops/pkg/logutil"
)

func TestInitlog(t *testing.T) {
	err := logutil.InitLog("/root/test.test")
	want := "open /root/test.test: permission denied"
	if want != err.Error() {
		t.Fatalf(`InitLog("/root/test.test") = %s, want "", error`, err.Error())
	}
}

func TestInfoLog(t *testing.T) {
	logutil.InitLog("test.test")
	logutil.InfoLog("info msg")
	logutil.ErrorLog("error msg")
	logutil.WarningLog("warning msg")
	data, _ := os.ReadFile("test.test")
	want := 179
	// delete tmp file
	defer os.Remove("test.test")
	fmt.Println(len(data))
	os.Stdout.Write(data)
	if want != len(data) {
		t.Fatalf(`%d`, len(data))
	}
}
