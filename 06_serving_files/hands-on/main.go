package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type","text/html;charset=UTF-8")
	   io.WriteString(wr,"foo ran")
}
func dog(wr http.ResponseWriter, req *http.Request) {
	   err := tpl.ExecuteTemplate(wr,"dog.gohtml" ,req)
	   if err != nil {
	   	log.Fatal(err)
	   }
}
func y(wr http.ResponseWriter, r *http.Request) {
	 if err := tpl.ExecuteTemplate(wr,"y.gohtml",nil); err != nil {
	 	log.Println(err)
	 }

}


var tpl *template.Template

func init() {
	 tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

  http.Handle("/", http.HandlerFunc(foo))
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/y/", y)

	http.ListenAndServe(":8080",nil)

}
