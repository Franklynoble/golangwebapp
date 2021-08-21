package main

import (
	"io"
	"net/http"
	"os"
)

func main(){
     http.HandleFunc("/",dog)
     http.HandleFunc("/toby.jpg",dogPic) // asking for image at the root level
     http.ListenAndServe(":8080",nil)
}

func dog(wr http.ResponseWriter, req *http.Request) {
  wr.Header().Set("Content-Type","text/html; charset=utf-8")
      //this would ask for an image
  io.WriteString(wr,`<img src="toby.jpg">`)
}

func dogPic(wr http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(wr, "file not found",4040)
		return
	}
	defer f.Close()
   io.Copy(wr,f) // would copy the from the file to our response writer
}