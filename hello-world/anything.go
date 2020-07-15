package main

import "fmt"

func main() {
	fmt.Println("Hey Evan. You'll be a star one day!")
	foo()
	fmt.Println("something more")

	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited.")
}

// control flow
// (1) sequence
// (2) loop: iterative
// (3) conditional
