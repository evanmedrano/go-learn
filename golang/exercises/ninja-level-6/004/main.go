package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi! My name is %v %v and I am %v.", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Evan",
		last:  "Medrano",
		age:   29,
	}

	p1.speak()
}
