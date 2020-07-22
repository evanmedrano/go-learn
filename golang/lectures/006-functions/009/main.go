package main

import "fmt"

func main() {
	// first () returns the inner func (func() int)
	// second () returns the inner func's return value (23)
	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 23
	}
}
