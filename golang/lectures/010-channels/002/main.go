package main

import "fmt"

func main() {
	// bi-directonal channel
	c := make(chan int, 2)

	c <- 43
	c <- 42

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)

	// send-only channel
	cs := make(chan<- int, 2)

	cs <- 43
	cs <- 42

	fmt.Println("------")
	fmt.Printf("%T\n", cs)

	// receive-only channel
	cr := make(<-chan int, 2)

	fmt.Println("------")
	fmt.Printf("%T\n", cr)
}
