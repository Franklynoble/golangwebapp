package main

import (
	"log"
	"os"
	"text/template"
)
type car struct {
	Manufacturer string
	Model string
	Doors int
}
type items struct {
	Wisdom [] sage
	Transport []car
}

var tpl *template.Template
type sage struct {
	Name string
	Motto string
}

func init() {
	 tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	b := sage {
		Name: "Buddha",
		Motto: "The Belief of no belief",
	}
	
	g := sage{
		Name:"Gandhi",
		Motto:"Be the Change",
	}
	ml := sage{
		Name:  "Martin Lurther king",
		Motto: "Hatred never ceases with love alone is healed hatred",
	}
	J := sage{
		Name:  "JesusChrist",
		Motto: "Loves All",
	}
	m := sage{
		Name:  "Muhammad",
		Motto: "to Overcome Evil with Good,To resiste evil by evil, is evil",
	}


	f := car {
		Manufacturer: "Ford",
		Model: "F150",
		Doors: 2,
	}
	c := car {
		Manufacturer: "Toyota",
		Model: "Corolla",
		Doors: 4,
	}
	sages1 := []sage{m,b,g,J,ml}
	cars := []car{f,c}

	data := items{
		Wisdom:    sages1,
		Transport: cars,
	}

	//sages := []sage{budha,gandhi,JesusChrist,mlk,muhammad}


	if err := tpl.Execute(os.Stdout,data); err != nil {
		log.Fatalln(err)
	}


}
