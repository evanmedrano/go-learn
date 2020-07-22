package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("The area for %T is %v\n", s, s.area())
}

func main() {
	c := circle{10}
	s := square{10}

	info(c)
	info(s)
}
