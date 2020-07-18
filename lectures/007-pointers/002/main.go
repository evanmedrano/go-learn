// Use pointers when you don't want to pass around a large chunk of data
// around in your program – pass the address where the data is stored
// instead (more performant)

// when you need to change a value at a particular location in memory
// (dereference the pointer and store the value at the address)

// changing the value at the same address as seen below, is mutation

// Method sets are the set of methods attached to a particular type
// non-pointer receivers work for both pointer and non-pointer values
// pointer receivers only work for pointer values
// (look over definition in Go docs)

package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 23
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
