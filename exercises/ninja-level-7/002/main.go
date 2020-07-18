package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "Evan"
	p.last = "Medrano"
}

func main() {
	p1 := person{
		first: "Jillian",
		last:  "Weimer",
	}

	fmt.Println(p1)
	fmt.Println(&p1.first, &p1.last)

	changeMe(&p1)

	fmt.Println(p1)
	fmt.Println(&p1.first, &p1.last)
}
