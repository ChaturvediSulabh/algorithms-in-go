package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(sum([]int{123, -456, 789}))
	fmt.Println(sum([]int{123, 45, 67, 89}))

	fmt.Println(recSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(recSum([]int{123, -456, 789}))
	fmt.Println(recSum([]int{123, 45, 67, 89}))
}

func sum(list []int) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func recSum(list []int) int {
	if len(list) == 0 {
		return 0
	} else if len(list) == 1 {
		return list[0]
	}
	return list[0] + recSum(list[1:])
}
