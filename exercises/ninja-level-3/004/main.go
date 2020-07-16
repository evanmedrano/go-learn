package main

import "fmt"

func main() {
	birthYear := 1991

	for {
		if birthYear > 2020 {
			break
		}

		fmt.Println(birthYear)
		birthYear++
	}
}
