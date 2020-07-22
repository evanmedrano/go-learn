package word

import "strings"

// UseCount returns the total amount a word was used in a string (s).
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the total word count for a given string (s).
func Count(s string) int {
	return len(strings.Fields(s))
}
