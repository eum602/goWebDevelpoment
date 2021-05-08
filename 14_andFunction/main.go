package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Adim  bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.goHtml"))
}

func main() {
	u1 := user{
		Name:  "eum602",
		Motto: "nothing",
		Adim:  true,
	}

	u2 := user{
		Name:  "Jhon Doe",
		Motto: "nothing",
		Adim:  false,
	}

	u3 := user{
		Name:  "",
		Motto: "nothing",
		Adim:  false,
	}

	users := []user{u1, u2, u3}
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
	//RUN WITH: go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
