package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	result := foo(nums...)

	fmt.Println(result)
}

func foo(x ...int) int {
	result := 0

	for _, v := range x {
		result += v
	}

	return result
}
