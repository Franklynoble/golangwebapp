package main




import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)   //convert the Variable  Value of string to int
	if err != nil {
		log.Fatalln(err)
	}
	count++   // increase the converted var
	cookie.Value = strconv.Itoa(count)  //convert back to string

	http.SetCookie(res, cookie)  // set the cookie again

	io.WriteString(res, cookie.Value)  // write to the browser
}