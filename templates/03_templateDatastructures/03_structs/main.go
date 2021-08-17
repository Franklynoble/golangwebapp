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

	if err := tpl.Execute(os.Stdout,budha); err != nil {
		log.Fatalln(err)
	}


}
