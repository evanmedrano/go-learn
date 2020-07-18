package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 7, 8}
	sum := foo(xi...)
	fmt.Println(sum)
	people := []string{"Evan", "Matt"}
	greeting("Hello", people...)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}

func greeting(prefix string, who ...string) {
	fmt.Println(prefix, ": ", who)
}
