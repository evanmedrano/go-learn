package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me takeo take care of that for you, James?",
			"I would really prefer to be across secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println("\nSorted by age", users)
	sort.Sort(ByLast(users))
	fmt.Println("\nSorted by last name", users)

	for _, v := range users {
		fmt.Println("\nSorted sayings for", v.First, v.Last)
		sort.Strings(v.Sayings)

		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}
}
