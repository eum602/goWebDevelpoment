package main

import (
	"io"
	"net/http"
)

type myType int

func (m myType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog/":
		io.WriteString(res, "something in /dog")
	case "/cat":
		io.WriteString(res, "something in /cat")
	}
}

func main() {
	var d myType
	/*
		"d" is:
		*  of type "myType"
			* which also implements the "Handler" interface
			* and has an underlying type of "int"
	*/
	http.ListenAndServe(":8080", d) //anything arriving to port 8080 is going to be handled by "d"

}
