package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar", req.Method)
	//set the new location where the client can get the resource
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther) //redirects with status code 303
	//what it does is send the client to "/"
}

func barred(w http.ResponseWriter, req *http.Request) {
	//if a user goes to "/barred" then the form will hit "/bar" which in turn will move with status "303 - see other" to "/"
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
