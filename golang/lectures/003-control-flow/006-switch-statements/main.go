package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print")
	case 4 == 4, 2 < 10:
		fmt.Println("prints")
		fallthrough
	case 8 == 8:
		fmt.Println("also prints")
		fallthrough
	case 2 == 1:
		fmt.Println("also prints, even though it's false")
		fallthrough
	default:
		fmt.Println("this is the default case")
	}
}
