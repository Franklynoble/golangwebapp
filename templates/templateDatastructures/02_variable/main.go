package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}


func main() {

	sages := []string{"Grandhi","MLK","Buddha","Jesus","Muhammad"}
    err := tpl.Execute(os.Stdout,sages)
    if err != nil {
    	log.Fatalln(err)
    }else {
    	fmt.Println()
    }






}
