package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.goHtml"))
}

func main() {
	g1 := struct {
		Score1 int
		Score2 int
	}{
		10,
		20,
	}
	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
	//RUN WITH: go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
