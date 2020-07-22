package main

import "fmt"

func main() {
	favSport := "Baseball"

	switch favSport {
	case "Basketball":
		fmt.Println("Not my favorite sport")
	case "Baseball":
		fmt.Println("My favorite sport is baseball")
	default:
		fmt.Println("I don't know that sport")
	}
}
