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
	var log Content
	InitLog("test.test")
	log = Content{"info msg"}
	log.InfoLog()
	log = Content{"error msg"}
	log.ErrorLog()
	log = Content{"warning msg"}
	log.WarningLog()
	data, _ := os.ReadFile("test.test")
	want := 116
	// delete tmp file
	defer os.Remove("test.test")
	fmt.Println(len(data))
	if want != len(data) {
		t.Fatalf(`%d`, len(data))
	}
}
