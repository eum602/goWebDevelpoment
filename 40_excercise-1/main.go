package main

import (
	"io"
	"net/http"
	"text/template"
)

type home int
type dog int
type me int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func (m home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "some message")
}

func (m me) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from me!!!")
}

func (m dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from dog!!!")
}

func main() {
	var d dog
	var m me
	var h home
	http.Handle("/", h)
	http.Handle("/dog/", d)
	http.Handle("/me/", m)
	http.ListenAndServe(":8080", nil)

}
