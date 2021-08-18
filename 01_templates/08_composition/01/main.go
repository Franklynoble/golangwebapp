package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template


type person struct {
	 Name string
	 Age int

}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := person{
		Name:"James Bond",
		 Age: 42,
	}



	if err := tpl.Execute(os.Stdout,p1); err != nil {
		log.Fatal(err)
	}
}
