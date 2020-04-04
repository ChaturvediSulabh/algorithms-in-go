package main

import "fmt"

func main() {
	fmt.Println(decimalToBase(10, 2))
	fmt.Println(decimalToBase(10, 8))
	fmt.Println(decimalToBase(10, 16))
	fmt.Println(decimalToBase(-10, 2))
	fmt.Println(decimalToBase(-10, 8))
	fmt.Println(decimalToBase(-1010, 16))
}

func decimalToBase(num, base int) string {
	switch base {
	case 2:
		return fmt.Sprintf("%b", num)
	case 8:
		return fmt.Sprintf("%o", num)
	case 16:
		return fmt.Sprintf("%X", num)
	}
	return ""
}
