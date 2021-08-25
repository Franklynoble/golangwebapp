package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.cio", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id := uuid.NewV4() // getting the cookies Id
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),

			//Secure:   true, using the https
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
