package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/download.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/download.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("download.jpg")

	if err != nil {
		//http.Error returns an error to the client with the specified status code.
		http.Error(w, "file not found", http.StatusNotFound) //means 404 -> https://pkg.go.dev/net/http -> check common http methods
		return
	}

	defer f.Close()

	fi, err := f.Stat()

	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound) //means 404
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
