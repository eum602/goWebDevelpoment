package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) //here it serves all the files into this current directory
	// http.Dir(".") ---> cast to a Dir type which implements the FileSystem interface since Dir implements the 'Open' method
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="download.jpg">`) //when the browser requests this file then the closest match is "/"
	//which serves that.
}

//served files
// /download.jpeg
// /main.go
