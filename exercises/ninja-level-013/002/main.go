package main

import (
	"fmt"
	"github.com/evanmedrano/learning/exercises/ninja-level-013/002/quote"
	"github.com/evanmedrano/learning/exercises/ninja-level-013/002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
