package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go fanOutIn(c)

	receive(c)

	fmt.Println("Finished!")
}

func fanOutIn(c chan int) {
	for i := 0; i < 10; i++ {
		func() {
			for j := 0; j < 10; j++ {
				c <- i
			}
		}()
	}
}

func receive(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}
