package main

import "fmt"

func main() {
	f := foo()
	f()
}

func foo() func() {
	return func() {
		fmt.Println("I am a returned func!")
	}
}
