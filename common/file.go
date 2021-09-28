package common

import (
	"os"
	"strings"
)

func init() {
	// fmt.Println("file mould init funcation")
}

// Addfile
func Addfile(name string, content string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// fmt.Println(err)
		IfErrPrintE(err.Error())
	}
	content = strings.Join([]string{content, "\n"}, "")
	f.WriteString(content)
	defer f.Close()
	// fmt.Printf("Write %v sucessfully \n", name)
}
