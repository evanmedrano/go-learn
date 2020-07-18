package main

import "fmt"

func main() {
	fmt.Println(foo()())
}

func foo() func() int {
	x := 23

	return func() int {
		return x * x
	}
}
