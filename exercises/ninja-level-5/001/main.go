package main

import "fmt"

type person struct {
	first                   string
	last                    string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		first:                   "Evan",
		last:                    "Medrano",
		favoriteIceCreamFlavors: []string{"Vanilla", "Coffee"},
	}

	p2 := person{
		first:                   "Jillian",
		last:                    "Weimer",
		favoriteIceCreamFlavors: []string{"Chocolate", "Moose Tracks"},
	}

	fmt.Println(p1.first, p1.last)

	for i, v := range p1.favoriteIceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first, p2.last)

	for i, v := range p2.favoriteIceCreamFlavors {
		fmt.Println(i, v)
	}
}
