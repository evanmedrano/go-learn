package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Starting")
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin goroutines", runtime.NumGoroutine())

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go foo(wg)
	go bar(wg)

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Done!")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end goroutines", runtime.NumGoroutine())
}

func foo(wg *sync.WaitGroup) {
	for i := 0; i < 15; i++ {
		fmt.Println("In foo(), currently on number", i)
	}
	wg.Done()
}

func bar(wg *sync.WaitGroup) {
	for i := 0; i < 15; i++ {
		fmt.Println("In bar(), currently on number", i)
	}
	wg.Done()
}
