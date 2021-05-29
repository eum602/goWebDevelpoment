package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor", http.FileServer(http.Dir("vendor"))))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
