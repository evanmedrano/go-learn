package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	c2 := make(chan int, 1)

	c2 <- 32

	fmt.Println(<-c2)
}
