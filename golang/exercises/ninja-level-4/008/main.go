package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "CS"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)

	for k, xs := range m {
		fmt.Println("This is the record for", k)

		for i, v := range xs {
			fmt.Println("\t", i, v)
		}
	}
}
