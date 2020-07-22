package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	c := circle{
		radius: 10.0,
	}

	s := square{
		length: 5.0,
		width:  5.0,
	}

	info(c)
	info(s)
}
