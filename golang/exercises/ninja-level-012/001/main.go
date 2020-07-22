package main

import (
	"fmt"
	"github.com/evanmedrano/learning/exercises/ninja-level-012/001/dog"
)

func main() {
	y := 29

	v := dog.Years(y)

	fmt.Printf("I am %d years old in dog years.", v)
}
