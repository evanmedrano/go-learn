package main

import "fmt"

func main() {
	v := average([]float64{1, 2})
	fmt.Println(v)
}

func average(f []float64) float64 {
	var total float64

	for i := 0; i < len(f); i++ {
		total += f[i]
	}

	return total / float64(len(f))
}
