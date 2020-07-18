package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am a secret agent", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am a person", p.first, p.last)
}

// a value can be more than 1 type
// any value that implements speak() is of type human
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Hello human", h.(person).first)
	case secretAgent:
		fmt.Println("Hello human", h.(secretAgent).first)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"Evan",
			"Medrano",
		},
		ltk: false,
	}

	p1 := person{
		"Jillian",
		"Weimer",
	}

	sa1.speak()
	bar(sa1)
	p1.speak()
	bar(p1)
}
