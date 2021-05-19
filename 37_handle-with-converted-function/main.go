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
	// http.HandlerFunc is a type that implements the "Handler" interface.
	//http.Handle is a method that requires as a second argument a type that implements the "Handler" interface
	//Since any variable that is http.HandlerFunc implements the "Handler" interface then it can be passed as an
	//argument for the "Handle" method.
	//http.HandlerFunc has as an underlying type "func(res http.ResponseWriter, req *http.Request)"
	//Also we have two methods, "c" and "d" with this signature: (res http.ResponseWriter, req *http.Request)
	//then  in order to cast "c" and "d" as type "HandlerFunc" we can do -> HandlerFunc(c) and HandlerFunc(d)
	http.Handle("/dog/", http.HandlerFunc(d)) //http.HandlerFunc(d) converts to a type HandlerFunc and thus passes a type "Handler"
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil) //passing nil means to use the DefaultServeMux
}
