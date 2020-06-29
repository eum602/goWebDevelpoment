package main

import(
	"text/template"
	"log"
	"os"
)

func main()  {
	tpl,err := 	template.ParseGlob("templates/*") //loads all templates
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout,nil) //loads the first
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout,"tpl2.goHtml",nil) //with ExecuteTemplate we have to specify which template to load
	if err != nil {
		log.Fatalln(err)
	}
}