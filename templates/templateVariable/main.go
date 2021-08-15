package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl =  template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	//we can also pass in aggregate data
	// using the dot notation to parse data and you can only parse one p piece of data
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",`release self-focus; embrace other-focus`)
	if err != nil {
		log.Fatalln(err)
	}
}

