package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "something in /dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "something in /cat")
}

func main() {
	//using http.handleFunc in order to set the handlers for the DefaultServeMux
	//HandleFunc requires:
	//	* a. pattern
	//	* b. a functions whose signature is -> func(res http.ResponseWriter, req *http.Request)
	http.HandleFunc("/dog/", d) //matches all cases like /dog, /dog/some, /dog/some/somethingelse
	http.HandleFunc("/cat", c)  //it is a exact match, thuse /cat is matched but /cat/abc isn't

	http.ListenAndServe(":8080", nil) //passing nil means to use the DefaultServeMux
}
