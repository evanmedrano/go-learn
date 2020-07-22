package main

import "fmt"

func main() {
	// blocked channel
	// c := make(chan int)

	// c <- 42

	// fmt.Println(<-c)

	// channel fixed
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// successful buffer
	ch := make(chan string, 1)

	ch <- "Evan Medrano"

	fmt.Println(<-ch)

	// unsuccessful buffer
	// buffer value != values added to channel
	// d := make(chan int, 1)

	// d <- 23
	// d <- 33

	// fmt.Println(<-d)
}
