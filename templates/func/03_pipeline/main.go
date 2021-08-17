package main

import (
	"log"
	"math"
	"os"
	"text/template"
)



var tpl *template.Template
func double(x int) int {
	return x + x
}
func square(x int) int {
	return int(math.Pow(float64(x), 2))
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl": double,
	"fsq":  square,
	"fsqrt": squareRoot,
}


func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	if err :=  tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",3);err != nil {
		log.Fatal(err)
	}


}
