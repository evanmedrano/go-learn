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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	j := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	m := sage{
		Name:  "Muhammad",
		Motto: "Good to beat evil is good",
	}

	t := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	sages := []sage{g, j, m}
	cars := []car{t, f}

	data := struct {
		Widsom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
