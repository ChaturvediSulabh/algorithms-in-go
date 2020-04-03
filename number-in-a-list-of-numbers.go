package main

import "fmt"

func main() {
	fmt.Println(exists(5, []int{1, 2, 3, 4, 6, 5}))
	fmt.Println(exists(5, []int{1, 2, 3, 4, 6}))
	fmt.Println(exists(-5, []int{1, 2, 3, 4, 6, -5}))
	fmt.Println(exists(-5, []int{1, 2, 3, 4, 6, 5}))
}

func exists(n int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == n {
			return true
		}
	}
	return false
}
