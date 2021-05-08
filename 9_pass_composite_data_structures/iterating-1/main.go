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
	sages := []string{"Ghandi", "MLK", "Budha", "Jesus", "Muhamad"} //slices created with composite literals where
	//a type (string) and the values {"Ghandi", etc} are specified
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
	//execute with go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
