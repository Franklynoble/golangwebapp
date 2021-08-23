package main

import (
	"io"
	"net/http"
)

func main() {

http.HandleFunc("/",dog)
http.HandleFunc("/toby.jpg",dogPic)
http.ListenAndServe(":8080",nil)
}

func dog(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(
		wr,`<img src="toby.jpg">`)
}
func dogPic(wr http.ResponseWriter, req *http.Request) {
	http.ServeFile(wr,req,"toby.jpg")
}
