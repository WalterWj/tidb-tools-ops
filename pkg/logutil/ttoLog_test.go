package logutil

import "testing"

func TestInfoLog(t *testing.T) {
	want := "Content from log package!"
	logMsg := InfoLog()
}
