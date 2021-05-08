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
	nf, err := os.Create("index.html") //creates a null file
	if err != nil {
		log.Println("error creating file", err)
	}

	defer nf.Close()

	err = tpl.Execute(nf, nil) //prints the content over the nf file
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("tpl_another_added.goHtml", "tpl2.some") //I can add any template into the container even
	//when I declared "tpl" with some other themplates at the beggining
	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.some", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.goHtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) //simply gets the first file loaded into the tpl templates container
	if err != nil {
		log.Fatalln(err)
	}

}
