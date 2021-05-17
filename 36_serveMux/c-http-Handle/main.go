package main

import (
	"io"
	"net/http"
)

type myType int

func (m myType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "something in /dog")
}

type myType2 int

func (m myType2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "something in /cat")
}

func main() {
	var d myType
	var c myType2
	//As "nil" has been passed to the "ListenAndServe" method then the DefaultServeMux is used; in order to set
	//the handlers for that default mux the "Handle" method is used for that setting.
	//Note http.Handle is a different method that "Handle" method attached to *ServeMux method
	http.Handle("/dog/", d) //matches all cases like /dog, /dog/some, /dog/some/somethingelse
	http.Handle("/cat", c)  //it is a exact match, thuse /cat is matched but /cat/abc isn't

	http.ListenAndServe(":8080", nil) //by putting "nil" we specify to use the default serveMux
}
