package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some valueeee",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w,"in chrome go to: dev tools / application /10_cookies")
}

//read, would check to see if there is cookie, if yes it would read,
func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
    if err != nil {
    	http.Error(w,err.Error(), http.StatusNoContent)
	    return
    }
    fmt.Fprintln(w,"YOUR COOKIE",c)
}
