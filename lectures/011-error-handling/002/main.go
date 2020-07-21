package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	_, err = os.Open("no-file.txt")
	if err != nil {
		log.Panicln("there was an error", err)
	}
}
