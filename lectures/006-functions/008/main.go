package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("My fav num is", x)
	}

	f(23)
}

func foo() {
	fmt.Println("hello from foo")
}
