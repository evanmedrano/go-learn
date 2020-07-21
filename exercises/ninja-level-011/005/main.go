package main

import "fmt"

func main() {
	v := average([]float64{1, 2})
	fmt.Println(v)
}

func average(f []float64) float64 {
	var total float64

	for _, v := range f {
		total += v
	}

	return total / float64(len(f))
}
