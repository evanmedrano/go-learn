package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := even(sum, nums...)
	fmt.Println(result)
}

func sum(x ...int) int {
	result := 0

	for _, v := range x {
		result += v
	}

	return result
}

func even(f func(x ...int) int, ii ...int) int {
	var s []int

	for _, v := range ii {
		if v%2 == 0 {
			s = append(s, v)
		}
	}

	return f(s...)
}
