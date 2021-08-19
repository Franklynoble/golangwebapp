package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog)ServeHTTP(wr http.ResponseWriter,req *http.Request) {
	switch req.URL.Path {

	case "/dog":
		io.WriteString(wr, "dose dose dose")
	case "/cat":
		io.WriteString(wr, "kat , ket kat")
	}
}
func main(){
	var d hotdog
	http.ListenAndServe(":8080",d)


}