package main

import "fmt"

func main() {
	s := [][]string{[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James"}}

	fmt.Println(s)

	for _, xs := range s {
		fmt.Println(xs)

		for _, v := range xs {
			fmt.Println(v)
		}
	}
}
