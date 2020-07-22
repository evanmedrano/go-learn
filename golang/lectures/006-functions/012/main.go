package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)
	n2 := loopFactorial(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	result := 1

	for ; n > 0; n-- {
		result *= n
	}

	return result
}
