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
	tpl = template.Must(template.ParseFiles("index.html.gohtml"))
}

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Method string
		URL    *url.URL
		//takes map with key string, return the slice of string
		Submissions map[string][]string
		Header      http.Header
	}{
		Method:      req.Method,
		URL:         req.URL,
		Submissions: req.Form,
		Header:      req.Header,
	}
	tpl.ExecuteTemplate(w, "index.html.gohtml", data)
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}
