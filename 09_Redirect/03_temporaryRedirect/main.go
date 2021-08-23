package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/barred",barred)

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func foo (wr http.ResponseWriter, req *http.Request) {
	fmt.Println("Your method from foo:", req.Method ,"\n\n")
}

func bar(wr http.ResponseWriter, req *http.Request) {
   fmt.Print("Your Method, at bar:", req.Method)
   //process form submission here
   http.Redirect(wr,req, "/",http.StatusTemporaryRedirect)

}
func barred(wr http.ResponseWriter , req *http.Request) {
	fmt.Println("Your method at barred: ", req.Method , "\n")
	 tpl.ExecuteTemplate(wr,"index.gohtml",nil)
}