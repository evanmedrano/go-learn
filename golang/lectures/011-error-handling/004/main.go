package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v, %v, %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, norgateMathError{"23.000 N", "31332.22 W", nme}
	}
	return 42, nil
}