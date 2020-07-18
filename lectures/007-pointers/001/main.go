package main

import "fmt"

func main() {
	a := 23
	fmt.Println(a)
	// prepend & to view the address in memory where a value is stored
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	// type value is a pointer to an int (*int)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	// prepend * to get the value stored at the address
	// when having the address
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 42

	fmt.Println(a)
}
