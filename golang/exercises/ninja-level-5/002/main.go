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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println("Values for", v.first, v.last)

		for i, flav := range v.favoriteIceCreamFlavors {
			fmt.Println(i, flav)
		}
	}
}
