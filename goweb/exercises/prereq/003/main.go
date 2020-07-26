package main

import "fmt"

type person struct {
	age  int
	name string
}

func (p person) speak() {
	fmt.Println("Hey! I'm", p.name)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("Hey! I'm", sa.name, "and I have a license to kill.")
}

type human interface {
	speak()
}

func iDeclare(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Evan Medrano",
		age:  29,
	}

	sa1 := secretAgent{
		person: person{
			name: "James Bond",
			age:  32,
		},
		ltk: true,
	}

	iDeclare(p1)
	iDeclare(sa1)
}
