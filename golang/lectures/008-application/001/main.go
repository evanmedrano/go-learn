// Marshall returns the JSON encoding of the value

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"FirstName,omitempty"`
	Last  string `json:"LastName,omitempty"`
	Age   int    `json:",omitempty"`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		Last: "Moneypenny",
		Age:  27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(bs))
}
