// Package fmt implements formatted I/O with functions analogous to C's
// printf and scanf. The format 'verbs' are drived from C's but are simpler

// The verbs

// General
// %v  - the value in a default format
//       when printing structs, the plus flag (%+v) adds field names
// %#v - a Go-syntax representation of the value
// %T  - a Go-syntax representation of the type of the value
// %%  - a literal percent sign; consumes no value

// Boolean
// %t  - the word true or false

// Integer
// %b  - base 2
// %c  - the character represented by the corresponding Unicode code point
// %d  - base 10
// %o  - base 8
// %q  - a single-quoted character literal safely escaped with Go syntax
// %x  - base 16, with lower-case letters for a-f
// %X  - base 16, with upper-case letters for A-F
// %U  - Unicode format: U+1234; same as "U+%04X"

// check for the rest of the documentation at godoc.org/fmt

package main

import "fmt"

var y = 52

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Printf("%v, %T\n", s, s)
}
