package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
type sage struct {
	Name string
	Motto string
}

func init() {
	 tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	budha := sage {
		Name: "Buddha",
		Motto: "The Belief of no belief",
	}
	
	gandhi := sage{
		Name:"Gandhi",
		Motto:"Be the Change",
	}
	mlk := sage{
		Name:  "Martin Lurther king",
		Motto: "Hatred never ceases with love alone is healed hatred",
	}
	JesusChrist := sage{
		Name:  "JesusChrist",
		Motto: "Loves All",
	}
	muhammad := sage{
		Name:  "Muhammad",
		Motto: "to Overcome Evil with Good,To resiste evil by evil, is evil",
	}

	sages := []sage{budha,gandhi,JesusChrist,mlk,muhammad}


	if err := tpl.Execute(os.Stdout,sages); err != nil {
		log.Fatalln(err)
	}


}
