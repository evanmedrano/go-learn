package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}

type test struct {
	age    int
	answer int
}

var tests = []test{
	test{0, 0},
	test{1, 7},
	test{2, 14},
	test{3, 21},
	test{4, 28},
	test{5, 35},
}

func TestYears(t *testing.T) {
	for _, test := range tests {
		v := Years(test.age)

		if v != test.answer {
			t.Errorf("expected %v, got: %v", test.answer, v)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	for _, test := range tests {
		v := Years(test.age)

		if v != test.answer {
			t.Errorf("expected %v, got: %v", test.answer, v)
		}
	}
}
