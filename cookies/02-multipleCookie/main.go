package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read/",read)
	http.HandleFunc("/abundance/",abundance)
	http.ListenAndServe(":8080",nil)

}

func set(wr http.ResponseWriter, req *http.Request) {
	http.SetCookie(wr, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(wr, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprint(wr, "in chrome go to dev tool/ application / cookies")
}
func read(wr http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprint(wr, "YOUR COOKIE 1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Fprintln(wr, "YOUR COOKIE 2 :", c2)
	}

	c3, err := req.Cookie("specific")

	if err != nil {
		log.Fatal(err)
	} else {
      fmt.Fprintln(wr, "YOUR COOKIE 3",c3)
	}

}


func abundance(wr http.ResponseWriter , req *http.Request) {
	  http.SetCookie(wr,&http.Cookie{
		  Name:       "general",
		  Value:      "Some other value about general things",

	  })
	  http.SetCookie(wr, &http.Cookie{
		  Name:       "specific",
		  Value:      "some other value about specific things",

	  })
	  fmt.Fprintln(wr, "COOKIES WRITTEN - CHECK YOUR BROWSER")
      fmt.Fprintln(wr,"IN CHROME GO TO: dev tools / application / cookies")
}