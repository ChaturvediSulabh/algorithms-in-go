package main

import "fmt"

func main() {
	fmt.Println(fizzbuzz(10))
	fmt.Println(fizzbuzz(1))
	fmt.Println(fizzbuzz(111111))
	fmt.Println(fizzbuzz(15))
}

func fizzbuzz(num int) string {
	if num%15 == 0 {
		return "fizzbuzz"
	} else if num%5 == 0 {
		return "buzz"
	} else if num%3 == 0 {
		return "fizz"
	}
	return ""
}
