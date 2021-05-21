package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/mydog", dogPic) //2. this is called by "1"
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src="/mydog">
	`)// 1. calls /mydog
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("download.jpg") //3. this is finally delivered
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
