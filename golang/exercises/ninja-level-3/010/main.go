package main

import "fmt"

func main() {
	fmt.Printf("%v should be true\n", true && true)
	fmt.Printf("%v should be false\n", true && false)
	fmt.Printf("%v should be true\n", true || true)
	fmt.Printf("%v should be true\n", true || false)
	fmt.Printf("%v should be false", !true)
}
