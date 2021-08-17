package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}
type semester struct {
	Term   string
	Courses []course //
}
type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CSCI-40", "Introduction to programming in Go", "4"},
				{"CSCI-40", "Introduction to web Programming With Go", "4"},
				{"CSCI-40", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{Number: "CSCI-5", Name: "Advanced Go", Units: "5"},
				{Number: "CSCI-190", Name: "Advanced Web programming With Go", Units: "5"},
				{Number: "CSCI-191", Name: "Advanced Mobile Apps With Go", Units: "5"},
			},
		},

	}

	if err := tpl.Execute(os.Stdout, y); err != nil {
		log.Fatal(err)
	}

}
