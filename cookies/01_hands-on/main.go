package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var count = 0

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/set", set)


	http.HandleFunc("/ret", ret)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(wr http.ResponseWriter, req *http.Request) {
	http.SetCookie(wr, &http.Cookie{
		Name:  "my-one",
		Value: "val1",
		Path:  "/ret",


	})
//	count++
	fmt.Fprintln(wr, "Cookie Written: check the Browser")
	fmt.Fprintln(wr, "OPEN DEV TOOLS: ", count)
	fmt.Println(count)

}


func ret(wr http.ResponseWriter, req *http.Request) {
	dc, err := req.Cookie("my-one")

	if err != nil {
		log.Fatal(err, "err cached")
	} else {
         count ++
		fmt.Println()
		fmt.Fprintln(wr, "number of time written increased", count, dc)


	}

	wr.Header().Set("Content-Type", "text/html;charset=UTF-8")


	fmt.Println(count)
	io.WriteString(wr, "number of COde written")

}
