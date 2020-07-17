package main

import "fmt"

func main() {
	// 001 Array
	// Length is apart of the array's type [3] vs [4] are 2 diff types
	// var x [5]int
	// fmt.Println(x)
	// x[3] = 42
	// fmt.Println(x)
	// fmt.Println(len(x))

	// 002 Slice - composite literal
	// a SLICE allows you to group together VALUES of the same TYPE
	// y := []int{4, 43, 23, 45, 21}
	// fmt.Println(y)

	// // 003 Slice - for range
	// z := []string{"evan", "jillian", "matt", "lulu"}
	// fmt.Println(len(z))
	// fmt.Println(z[0])

	// for i, v := range z {
	// 	fmt.Println(i, v)
	// }

	// // 004 - Slicing a slice
	// x := []int{4, 5, 6, 1, 23}
	// fmt.Println(x[:1])

	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(x[i])
	// }

	// 005 - Appending to a slice
	// x := []int{4, 5, 6, 1, 23}
	// fmt.Println(x)

	// x = append(x, 33, 1000)
	// fmt.Println(x)

	// y := []int{21, 432, 55311, 123, 13}
	// x = append(x, y...)
	// fmt.Println(x)

	// 006 - Removing from a slice
	// x := []int{4, 5, 6, 1, 23}
	// y := []int{21, 432, 55311, 123, 13}
	// x = append(x, y...)
	// fmt.Println(x)

	// x = append(x[:2], x[4:]...)
	// fmt.Println(x)

	// 007 - Slice using make
	// When a slice's capacity is reached, the underlying array is replaced,
	// thus doubling the capacity
	// x := make([]int, 10, 100)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// 008 - mutli-dimensional slice
	// jb := []string{"James", "Bond"}
	// fmt.Println(jb)
	// em := []string{"Evan", "Medrano"}
	// fmt.Println(em)

	// xp := [][]string{jb, em}
	// fmt.Println(xp)

	// 009 - Map introduction
	// m := map[string]int{
	// 	"James":           31,
	// 	"Miss Moneypenny": 23,
	// }

	// fmt.Println(m)
	// fmt.Println(m["James"])
	// // If a key is not found in a map, the value returned is 0
	// fmt.Println(m["Evan"])

	// // checking if the value is actually in the map
	// v, ok := m["Evan"]
	// fmt.Println(v, ok)

	// if v, ok := m["James"]; ok {
	// 	fmt.Println("This is the if print", v)
	// }

	// 010 - Map - add element and range
	// m := map[string]int{
	// 	"James":           31,
	// 	"Miss Moneypenny": 23,
	// }

	// m["evan"] = 29

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// // slice of int (xi)
	// xi := []int{231, 1, 32, 4, 23}

	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }

	// 011 - Map - delete
	m := map[string]int{
		"James":           31,
		"Miss Moneypenny": 23,
	}
	m["evan"] = 29
	fmt.Println(m)
	delete(m, "evan")
	fmt.Println(m)
}
