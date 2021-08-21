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
    if err, _ :=	io.WriteString(wr,`<img src="toby.jpg">`); err != nil {

    }
}

func dogPic(wr http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(wr, "file not found",4040)
		return
	}
	defer f.Close()
	//io.Copy(wr,f) // would copy the from the file to our response writer

	fi, err := f.Stat()
	  if err != nil {
	  	http.Error(wr,"file not founc ",404)
		  return
	  }
	  //this is sparingly used, but just for educational purpose
	  http.ServeContent(wr,req,f.Name(),fi.ModTime(),f)
}