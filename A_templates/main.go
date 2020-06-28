package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//A template is like a container holding all the files
	tpl, err := template.ParseFiles("tpl.goHtml") //can receive many arguments to parse (variadic)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
