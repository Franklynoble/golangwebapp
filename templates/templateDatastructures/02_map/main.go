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

	sages := map[string]string{
	"India": "Gandhi",
	 "America":"Milk",
	  "Mediate":"Budha",
	   "Love": "Jesus",
	   "Prohpet":"Muhammed",
	}

 if 	err := tpl.Execute(os.Stdout,sages); err != nil {
 	     log.Fatalln(err)
 }



}
