package main

import "fmt"

// Use an anoymous struct when you need the data for a small area

func main() {
	p1 := struct {
		first, last string
		age         int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
