package main

import (
	"fmt"
	"net/http"
	)



func main() {
	http.HandleFunc("/",dog)
	//make sure you include this in your code if you are not handling your faveIcon
	http.Handle("/favicon.cio",http.NotFoundHandler())
    http.ListenAndServe(":8080",nil)
}


func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(w,"Go look at your terminal  ")
}
