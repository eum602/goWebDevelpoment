package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	//anything coming from a /resources and beyond that is "/resources/" is caught will fall in this handler
	//StripPrefix is gonna strip out "/resources"
	//then it is gonna take the second argument - which is a Handler - and use it to serve the required file served
	// from the specified folder , which in the example is "./assets"
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/download.jpg">`)
}

//to test it go to http://localhost:8080/
