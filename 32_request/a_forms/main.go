package main

import (
	"html/template"
	"log"
	"net/http"
)

type myType int

func (m myType) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() //necessary to do in order to make the data available
	if err != nil {
		log.Fatalln(err)
	}

	//here simply passing the req.forms as variables to be consumed by the template
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form) //req.Form -> is of type url.Values which in turn is a map
	//of string as key and a slice as a value
	//req.Form -> extracts data from the form and the url
	//req.PostForm extracts data ONLY from the form (which comes in the body or payload)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))

}

func main() {
	var d myType
	http.ListenAndServe(":8080", d)
}
