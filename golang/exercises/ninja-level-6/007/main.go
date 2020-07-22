package main

import "fmt"

func main() {
	f := foo
	f()
}

func foo() {
	fmt.Println("I am a func expression!")
}
