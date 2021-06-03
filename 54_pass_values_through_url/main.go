package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("abc") //searches for the "abc" parameter, if it si not present then "v" is null
	io.WriteString(w, "You wrote: "+v+" as a parameter")
}

//http://localhost:8080?abc=dog
