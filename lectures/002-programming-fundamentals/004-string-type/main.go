package main

import "fmt"

func main() {
	s := "Hello world!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var x string = "Evan Medrano"
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	bs := []byte(x)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
