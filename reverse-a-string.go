package main

import (
	"fmt"
)

func main() {
	fmt.Println(revStr("sulabh chaturvedi"))
}

func revStr(str string) string {
	newStr := ""
	for i := len(str) - 1; i > -1; i-- {
		newStr += string(str[i])
	}
	return newStr
}
