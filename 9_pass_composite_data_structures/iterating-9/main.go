package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.goHtml"))
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

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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

	data := struct { //using a composite literal instead of creating a new type. This is like an anonymous declaration type
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.goHtml", data)
	if err != nil {
		log.Fatalln(err)
	}
	//execute with go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
