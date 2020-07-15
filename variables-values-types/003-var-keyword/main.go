package main

import "fmt"

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functins, interfaces, slices, channels, and maps
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	// DELCARE the variable "y"
	// ASSIGN the value 43
	// delcare & assign = initialization
	var y = 43
	fmt.Println(y)

	fmt.Println(z)
}
