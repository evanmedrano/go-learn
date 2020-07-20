package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Hey! My name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Evan Medrano",
		age:  29,
	}

	// Works because a pointer is being passed into the saySomething func
	saySomething(&p1)

	// Does NOT work because the value is of type person and not &person
	// saySomething(p1)

	// However, this does work (even though it is not a pointer)
	p1.speak()
}
