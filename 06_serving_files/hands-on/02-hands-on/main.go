package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

//func dog(wr http.ResponseWriter, req *http.Request) {
//	wr.Header().Set("Content-Type","text/html;charset=UTF-8")
//	//   err := tpl.ExecuteTemplate(wr,"dog.gohtml" ,req)
//	       io.WriteString()
//	   if err != nil {
//	   	log.Fatal(err)
//	   }



func dogF(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type","text/html;charset=UTF-8")
	  io.WriteString(wr,`<img src="toby.jpg">`)

}
func dogPic(wr http.ResponseWriter, req *http.Request) {
	   f,err := os.Open("toby.jpg")

	   if err != nil {
	   	log.Println(err)
		   return
	   }
	   defer f.Close()

	     io.Copy(wr,f)
}


//
//var tpl *template.Template
//
//func init() {
//	 tpl = template.Must(template.ParseGlob("templates/*"))
//}

func main() {

  //http.Handle("/", http.HandlerFunc(foo))
	http.HandleFunc("/",dogF)
	http.HandleFunc("/dog/",dogPic)


	http.ListenAndServe(":8080",nil)

}
