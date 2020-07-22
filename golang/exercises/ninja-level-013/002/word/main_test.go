package word

import (
	"fmt"
	"github.com/evanmedrano/learning/exercises/ninja-level-013/002/quote"
	"reflect"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("This is my test string"))
	// Output:
	// 5
}

func ExampleUseCount() {
	fmt.Println(UseCount("This is my test string"))
	// Output:
	// map[This:1 is:1 my:1 string:1 test:1]
}

func TestCount(t *testing.T) {
	c := Count("This is my test string")
	if c != 5 {
		t.Error("Expected 5, got:", c)
	}
}

func TestUseCount(t *testing.T) {
	s := "This is my test string. I love to test things."
	got := UseCount(s)
	want := map[string]int{
		"This":    1,
		"is":      1,
		"my":      1,
		"test":    2,
		"string.": 1,
		"I":       1,
		"love":    1,
		"to":      1,
		"things.": 1,
	}
	eq := reflect.DeepEqual(got, want)

	if !eq {
		t.Errorf("Expected %v, got: %v", want, got)
	}
}
