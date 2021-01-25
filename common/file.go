package common

import (
	"os"
	"strings"
)

// Addfile
func Addfile(content string) {
	f, _ := os.OpenFile("users.sql", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	content = strings.Join([]string{content, "\n"}, "")
	f.WriteString(content)
	defer f.Close()
	// fmt.Printf("Write %v sucessfully \n", content)
}
