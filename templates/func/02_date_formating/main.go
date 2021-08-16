package main

import "html/template"

func init () {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles(""))
}
func main() {

}
