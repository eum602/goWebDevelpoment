package main

import (
	"net/http"
	"text/template"

	"github.com/satori/go.uuid"
)

var tpl *template.Template
var dbUsers = map[string]user{}      //userId -> userdata. Empty at the beginning
var dbSessions = map[string]string{} //using a composite literal, sessionId -> userId. Empty at the beginning

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct {
	UserName string
	First    string
	Last     string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	//http.Handle("favicon.ico", http.NotFoundHandler()) //conflicts with the handler of "/"
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sId, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(res, c)
	}

	//if the user already exists then get the user
	var u user
	if un, ok := dbSessions[c.Value]; ok { //";ok" means 'if true'
		u = dbUsers[un]
	}

	//process the submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u := user{un, f, l} //using a composite literal to create a new object
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(res, "index.gohtml", u)

}

func bar(res http.ResponseWriter, req *http.Request) {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
