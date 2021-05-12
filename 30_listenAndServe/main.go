package main

import (
	"fmt"
	"net/http"
)

type someType int

func (m someType) ServeHTTP(w http.ResponseWriter, r *http.Request) { //any "type" with ServeHttp that has an ResponseWriter
	//and a pointer to a Request is a HANDLER
	fmt.Fprintln(w, "Some code ...")

}

func main() {
	var a someType
	http.ListenAndServe(":8080", a) //listenAndServe method receives a port where to listen and a variable
	//that implements the ServeHTTP interface

}
