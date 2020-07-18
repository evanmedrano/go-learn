package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("anonymous func ran")
	}()
	func(x int) {
		fmt.Println("My fav num is", x)
	}(23)
}

func foo() {
	fmt.Println("hello from foo")
}
