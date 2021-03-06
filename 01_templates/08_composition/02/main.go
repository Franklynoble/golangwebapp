package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}
type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	p1 := doubleZero{
		person: person{
			Name: "Ian Fleming",
			Age:  56,
		},
		LicenseToKill: false,
	}

	if err := tpl.Execute(os.Stdout, p1); err != nil {
		log.Fatal(err)
	}

}
