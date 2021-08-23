package main
import (
	"html/template"
	"log"
	"net/http"
)

/*

hands On exercise using Default Mux, and adding convertion
type must be http.resWr , req http.req
*/


func dog(wr http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(wr, "dog.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}
func me(wr http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(wr, "me.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func index(wr http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(wr, "index.html.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	// tpl.Execute(wr,req)

}

var tpl *template.Template

func init() {
	//ParseGlobEnsures only one file is passed.
	 //all th files would be linked under this directory
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	//using the handler func to do conversion, this only accept http.resWr and  http.req type
	http.HandleFunc("/me/", http.HandlerFunc(me))
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
