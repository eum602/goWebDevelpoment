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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type sage struct {
	Name  string
	Motto string
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func main() {
	budha := sage{
		Name:  "Budha",
		Motto: "Something else",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	nlk := sage{
		Name:  "Martin Luther King",
		Motto: "some",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{
		budha,
		gandhi,
		nlk,
	}

	cars := []car{
		f,
		c,
	}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
	//execute with go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
