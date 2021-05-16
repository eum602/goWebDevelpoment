package main

import (
	"fmt"
	"net/http"
)

type someType int

func (m someType) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("some-key", "a token maybe")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> some code goes here</h1>")
}

func main() {
	var d someType
	http.ListenAndServe(":8080", d)
}
