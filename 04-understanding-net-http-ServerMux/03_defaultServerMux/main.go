package main

import (
	"io"
	"net/http"
)



//this is more elegant code to work with,using mux
func d(wr http.ResponseWriter, req *http.Request) {
	io.WriteString(wr,"Dogs domain Beware!")
}


func c (wr http.ResponseWriter, req *http.Request) {
	io.WriteString(wr,"Cat Cat cat ")
}
func main() {

	//create a new serveMux
	http.HandleFunc("/dog/",d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080",nil)

}
