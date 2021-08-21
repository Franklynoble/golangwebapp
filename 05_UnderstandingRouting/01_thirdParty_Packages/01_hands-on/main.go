package main
import (
	"html/template"
	"log"
	"net/http"
)

/*

hands On exercise using Default Mux
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
	err := tpl.ExecuteTemplate(wr, "index.gohtml", nil)
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

	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/", me)

	http.ListenAndServe(":8080", nil)
}
