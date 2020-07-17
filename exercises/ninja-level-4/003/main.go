package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}
