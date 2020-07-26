package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Address, City, Name, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Region:  "Northern",
			Address: "123 St",
			City:    "Los Angeles",
			Name:    "Monopoly Hotel",
			Zip:     "12345",
		},
		hotel{
			Region:  "Northern",
			Address: "Sunset Blvd",
			City:    "Los Angeles",
			Name:    "Fancy Hotel",
			Zip:     "90021",
		},
		hotel{
			Region:  "Central",
			Address: "456 Blvd",
			City:    "San Diego",
			Name:    "Shoddy Inn",
			Zip:     "90210",
		},
		hotel{
			Region:  "Central",
			Address: "Too Expensive Rd",
			City:    "Beverly Hills",
			Name:    "Luxury Hotel",
			Zip:     "90210",
		},
		hotel{
			Region:  "Southern",
			Address: "ABC Dr",
			City:    "San Fransico",
			Name:    "Hotel By the Bay",
			Zip:     "90024",
		},
		hotel{
			Region:  "Southern",
			Address: "Hollywood Blvd",
			City:    "Hollywood",
			Name:    "Hollywood Hotel",
			Zip:     "90120",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
