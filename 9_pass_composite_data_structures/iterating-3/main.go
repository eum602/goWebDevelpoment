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
	sages := map[string]string{
		"India": "Dheli",
		"Peru":  "Lima",
		"EEUU":  "Washington",
		"Japan": "Tokyo",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
	//execute with go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
