package main

import (
	"html/template"
	"log"
	"os"
	"time"
)
var tpl *template.Template

func init () {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now()); err != nil {
		log.Fatal(err)

	}
}
//this func get the value of type time
func monthDayYear(t time.Time)string {
	return t.Format("01-02-2006") // the format method called on value type time
}
 // use the template to get the funcMap to call the function
var fm = template.FuncMap{  // pass the method on func map
	      "fdateMDY":monthDayYear,  // the method is the value of Key "   "fdateMDY"
}
