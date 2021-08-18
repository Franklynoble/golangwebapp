package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct{
	Name  string
	Motto  string
	Admin  bool
}


func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user {
		Name: "Buddha",
		Motto: "The belief of no beliefs",
		 Admin: false,
	}
u2 := user{
	Name:  "Gandhi",
	Motto: "Be the change",
	Admin: true,
}

u3 := user{
	Name:  "",
	Motto: "Nobody",
	Admin: true,
}
users := []user{u1,u2,u3}

if err := tpl.Execute(os.Stdout,users); err != nil {
	log.Fatal(err)
}


}