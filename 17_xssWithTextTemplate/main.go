package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Page struct {
	Title   string
	Heading string
	Input   string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing else added",
		Input:   `<script>alert("hello triggered from html")</script>`, //this is a vulnerability
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}

//go run main.go > index.html &&  chromium-browser index.html && rm index.html
