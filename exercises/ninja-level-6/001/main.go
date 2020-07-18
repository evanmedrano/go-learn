package main

import "fmt"

func main() {
	f := foo()
	s, i := bar()
	fmt.Println(f, s, i)
}

func foo() int {
	return 23
}

func bar() (string, int) {
	return "Evan Medrano", 23
}
