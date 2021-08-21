package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/",dog)

	http.ListenAndServe(":8080",nil)

}

func dog(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(wr, `
<!--image does not serve, we are requesting this file from our server but the file isn been set up to serve yet; on our server-->
  
<img src="/toby.jpg">`)


}
