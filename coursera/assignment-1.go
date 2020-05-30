package main

import (
	"fmt"
	"log"
)

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	fmt.Println(sum(num1, num2))
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func isError(err error) {
	if err != nil {
		log.Fatal()
	}
}
