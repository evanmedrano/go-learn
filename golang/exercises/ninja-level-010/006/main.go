package main

import (
	"fmt"
)

func main() {
	c := populate()

	receive(c)

	fmt.Println("finshed!")
}

func populate() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
