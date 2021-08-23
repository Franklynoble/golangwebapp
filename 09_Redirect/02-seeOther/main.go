package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/barred",barred)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)



}
func foo(wr http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo:",req.Method, "\n\n")
}
//when bar is access,it would automatically be redirected to foo
func bar(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar ",req.Method, "\n\n")
    //process form submission here
       wr.Header().Set("Location","/")  // the response header url location to redirect to
	wr.WriteHeader(http.StatusSeeOther)
}

func barred(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ",req.Method)
	tpl.ExecuteTemplate(wr,"index.gohtml",nil)
}