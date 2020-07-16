package main

import "fmt"

func main() {
	// can also use %c which is the character represented by the
	// cooresponding Unicode code point

	for i := 33; i <= 122; i++ {
		fmt.Printf("Char coorespending to Ascii Code %d = %#U\n", i, i)
	}
}
