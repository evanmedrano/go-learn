package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Error was found: %v", c.info)
}

func foo(e error) {
	// assertion is used here to make sure the error is of type customErr
	fmt.Println(e.(customErr))
}

func main() {
	c := customErr{
		info: "I love coding",
	}
	foo(c)
}
