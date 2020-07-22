package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 1000})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3}))
	// Output:
	// 2
}

type test struct {
	values []int
	answer float64
}

var tests = []test{
	test{[]int{1, 2, 3}, 2},
	test{[]int{10, 2, 3, 1, 0}, 2},
	test{[]int{1, 2, 3, 3, 10, 2, 4}, 2.8},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range tests {
		ca := CenteredAvg(v.values)

		if ca != v.answer {
			t.Errorf("Expected %v, got: %v", ca, v)
		}
	}
}
