package main

import "fmt"

func main() {
	evan := struct {
		first string
		last  string
		age   int
	}{
		first: "Evan",
		last:  "Medrano",
		age:   29,
	}

	fmt.Println(evan)
}
