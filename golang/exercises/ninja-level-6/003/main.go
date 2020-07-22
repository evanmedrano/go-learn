package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println(", and I will go second!")
}

func bar() {
	fmt.Println("I will go first")
}
