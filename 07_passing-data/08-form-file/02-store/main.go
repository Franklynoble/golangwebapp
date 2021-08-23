package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)


/*
Note! this code an empty file must always be created as a place holder for this file
to be read,and with its extension
*/
var tpl  *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	//if the request method is post, this code runs
	if req.Method == http.MethodPost {
		//open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

	}
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	 tpl.ExecuteTemplate(w,"index.gohtml",s)
}
