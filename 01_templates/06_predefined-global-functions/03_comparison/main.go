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

	g1 := struct{
		Score1 int
		Score2 int

	}{
		Score1: 7,
		Score2: 9,
	}

	if err := tpl.Execute(os.Stdout,g1); err != nil{
		log.Fatal(err)
	}

}
