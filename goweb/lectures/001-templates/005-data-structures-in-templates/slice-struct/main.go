package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but love alone is healed",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good",
	}

	sages := []sage{buddha, jesus, mlk, gandhi, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
