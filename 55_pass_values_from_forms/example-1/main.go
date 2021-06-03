package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//if the form is sent as a "post" -> the info travels in the request body,
	//otherwise if the message is sent as
	//as a "get" then the information travels in the "url"
	io.WriteString(w, `
	<form method="get">
		<input type="text" name=q>
		<input type="submit">
	</form>
	<br>`+v)
}

// http://localhost:8080/foo
