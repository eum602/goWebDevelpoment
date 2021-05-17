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
	mux := http.NewServeMux() //returns *ServeMux which implements the method ServeHttp thus implements the Handler interface
	/*
		* *ServeMux also implements the methods
			* Handle -> takes a "pattern" and a "Handler" interface
			* HandleFunc
	*/
	//since "d" and "c" are types that implements the "Handler" interface then the can be passed to mux.Handle
	mux.Handle("/dog/", d) //matches all cases like /dog, /dog/some, /dog/some/somethingelse
	mux.Handle("/cat", c)  //it is a exact match, thuse /cat is matched but /cat/abc isn't

	http.ListenAndServe(":8080", mux) //since mux is of type *ServeMux and this implements the "Handler"
	// interface then it can be passed as the second argument in the "ListenAndServe" method
}

/*
Interchangeable terms:
	* router
	* request router
	* multiplexer
	* mux
	* servemux
	* server
	* http router
	* http request router
	* http multiplexer
	* http mux
	* http servemux
	* http server
*/