package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if  err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w,"index.html.gohtml",req.Form)

}

 var tpl *template.Template

  func init() {
  	tpl = template.Must(template.ParseFiles("index.html.gohtml"))
  }
func main() {

	var d hotdog
	http.ListenAndServe(":8080", d)
}

