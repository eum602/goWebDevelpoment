package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to CS -1", "4"},
				course{"CSCI-40", "CS-1", "4"},
				course{"CSCI-40", "AI", "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"A1-140", "Electric Circuits", "3"},
				course{"A1-141", "Neural Networks", "2"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{"A2-340", "Artificial Intelligence", "5"},
				course{"A1-341", "Robots, Analysis and Control", "2"},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}
}

//go run main.go > index.html &&  chromium-browser index.html && rm index.html
