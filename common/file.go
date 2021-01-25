package common

import (
	"os"
	"strings"
)

func addfile(context string) {
	f, _ := os.OpenFile("users.sql", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	context = strings.Join([]string{context, "\n"}, "")
	f.WriteString(context)
	defer f.Close()
	// fmt.Printf("Write %v sucessfully \n", context)
}
