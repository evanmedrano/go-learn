package main

import "fmt"

type person struct {
	age  int
	name string
}

func (p person) pSpeak() {
	fmt.Println("Hey! I'm", p.name)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) saSpeak() {
	fmt.Println("Hey! I'm", sa.name, "and I have a license to kill.")
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

	fmt.Println(p1.name)
	p1.pSpeak()

	fmt.Println(sa1.name)
	sa1.saSpeak()
	sa1.pSpeak()
}
