package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
func init () {
	tpl = template.Must(template.ParseGlob("01_templates/*.gohtml"))

}
func main () {

	if err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml",42); err != nil {
		log.Fatal(err)
	}
}

