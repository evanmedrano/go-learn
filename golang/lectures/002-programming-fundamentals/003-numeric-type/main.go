// int and float64 are the defaulted types given by the compiler

// unsigned int values (uint8) start at 0
// regular int values (int8) include negative values

// all numeric types are distinct except
// byte which is alias for uint8
// rune which is an alias for int32

package main

import "fmt"

var x int
var y float64

func main() {
	x = 42
	y = 42.34532
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
