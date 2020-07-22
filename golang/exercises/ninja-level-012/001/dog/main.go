// Package dog handles all dog-related conversions.
package dog

// Years converts human years to dog years (1 human year == 7 dog years)
func Years(n int) int {
	return n * 7
}
