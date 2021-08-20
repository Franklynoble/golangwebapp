package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var tpl *template.Template

func init() { 
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		Method:      req.Method,
		URL:         req.URL,
		Submissions: req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}
