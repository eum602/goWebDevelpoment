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

type sage struct {
	Name  string
	Motto string
}

func main() {
	budha := sage{
		Name:  "Budha",
		Motto: "Something else",
	}
	err := tpl.Execute(os.Stdout, budha)
	if err != nil {
		log.Fatalln(err)
	}
	//execute with go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
