package main

import (
	"io"
	"net/http"
)

type hotdog int

func (c hotdog) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	io.WriteString(wr,"Dogs domain Beware!")
}

type hotcat int
func (c hotcat) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	io.WriteString(wr,"Cat Cat cat ")
}
func main() {
	//create values for type hotdog,hotcat
var d  hotdog
var c hotcat
//create a new serveMux
 mux := http.NewServeMux()
 mux.Handle("/dog/",d)
 mux.Handle("/cat",c)
 http.ListenAndServe(":8080",mux)

}
