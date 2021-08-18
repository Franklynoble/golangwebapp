package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}
func main() {

	xs := []string{"zero","one","two","Three","four","five"}

	data := struct {
		Words []string
		lName string
	}{
		Words: xs,
		lName: "Noble",
	}
	if err := tpl.Execute(os.Stdout,data); err != nil {
		log.Fatal(err)
	}

}
