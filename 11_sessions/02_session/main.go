package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user Id, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/bar",bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
    http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//get cookie
	c, err := req.Cookie("session") //ask for a cookie if cookie is not found, create one below

	if err != nil {
		sID := uuid.NewV4() // using the uuid as a value for the Cookie
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c) // set the cookie
	}

	//if the user exists already, get user
	var u user

	if un, ok := dbSessions[c.Value]; ok { //c.Value is where we have stored our
		// unique session ID, get the value on the Cookie, the cookie value is the key in this Map
		u = dbUsers[un]
	}

	//process form submission

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		i := req.FormValue("lastname")
		u = user{
			un, f, i,	}
		// if dbSessions c.Value key is used to get the value un which is
		  dbSessions[c.Value] = un  //use the c.Value as key to get the  username,
		  dbUsers[un] = u //we use the value of, username as the key of map dbUsers
		                   //using this key, you get all the value of users
	}
	tpl.ExecuteTemplate(w,"index.gohtml",u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get the Cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w,req, "/",http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w,req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w,"bar.gohtml",u)
}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV