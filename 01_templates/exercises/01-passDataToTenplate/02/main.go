package main

import (
	"html/template"
	"log"
	"os"
)


type address struct {
	Street string
	Name   string
	City   region
	Zip    string
}
type region struct {
	Southern, Central, Northern string
}

type california struct {
	hotels


}
type hotels struct {
	 Names  string
	Addresses []address
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
   //  n := map[string] interface{}{}
	 cali := []hotels{
		hotels {
			 Names: "GtView",
			 Addresses: []address{{Street: "no.302 sow mill farms, sidney",Name: "GtView", City: region{
				 Central: "Central Region",
			 }, Zip: "2001",
			 },
			 },
		 },
		 hotels{
			 Names:     "JazzFresi",
			 Addresses: []address{{ Street: "no 102. wilbur wright Street,new hil top",
				 Name: "JazzFresi Hotels",
				 City: region{
					 Southern: "Southern Region",
				 },
				 Zip:  "10022",
			 }},
		 },
		 hotels{

			 Names:     "SmartView",
			 Addresses: []address{{
			 	Street: "no 40 Mcm Fredrick Street, north park",
				 Name:  "SmartView Hotels",
				 City: region{
					 Southern: "Northern Region",
				 },
				 Zip:  "1001",
			 }},
		 },



	 }


	if err := tpl.Execute(os.Stdout, cali); err != nil {
		log.Fatal(err)
	}

}
