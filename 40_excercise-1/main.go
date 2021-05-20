package main

import (
	"io"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "some message")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from me!!!")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from dog!!!")
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
