package main

import (
	"io"
	"net/http"
)

type home int
type dog int
type me int

func (m home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from home!!!")
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
