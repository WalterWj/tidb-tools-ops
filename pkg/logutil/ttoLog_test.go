package logutil

import (
	"fmt"
	"os"
	"testing"
)

func TestInitlog(t *testing.T) {
	err := InitLog("/root/test.test")
	want := "open /root/test.test: permission denied"
	if want != err.Error() {
		t.Fatalf(`InitLog("/root/test.test") = %s, want "", error`, err.Error())
	}
}

func TestInfoLog(t *testing.T) {
	InitLog("test.test")
	InfoLog("info msg")
	ErrorLog("error msg")
	WarningLog("warning msg")
	data, _ := os.ReadFile("test.test")
	want := 116
	// delete tmp file
	defer os.Remove("test.test")
	fmt.Println(len(data))
	if want != len(data) {
		t.Fatalf(`%d`, len(data))
	}
}
