package main

import (
	"html/template"
	"log"
	"os"
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
		Input:   `<script>alert("hello triggered from html")</script>`, //this is NO a vulnerability anymore
		//the html/template detects it and modifies the outputed html in order to not to trigger potential non desired scripts
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}

//go run main.go > index.html &&  chromium-browser index.html && rm index.html
