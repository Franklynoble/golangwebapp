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

	xs := []string{"one","Two","Three","Four","Five"}
	if err := tpl.Execute(os.Stdout,xs); err != nil {
		log.Fatal(err)
	}
}
