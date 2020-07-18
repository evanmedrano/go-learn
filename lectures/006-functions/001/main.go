package main

import "fmt"

func main() {
	foo()
	bar("Evan")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Evan", "Medrano")
	fmt.Println(x, y)
}

// func (r receiver) identifier(parameters) (returns(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
