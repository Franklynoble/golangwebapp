package main

import (
	"io"
	"net/http"
)

func main() {
http.HandleFunc("/", dog)
//this would go to the directory, assets. in the assets folder, it would display the toby.jpg.
//Note. the Directory,
//it would totally strip off the link resources
http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("./assets"))))
 http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/html; charset=utf-8")
   io.WriteString(w,`<img src="/resources/toby.jpg">`)
}
