package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 23)
	if err != nil {
		log.Fatalln(err)
	}
}