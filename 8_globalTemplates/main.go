package main

import(
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*")) //Template.Must CHECKS FOR ERRORS
}

func main()  {	
	err := tpl.Execute(os.Stdout,nil) //loads the first
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout,"tpl2.goHtml",nil) //with ExecuteTemplate we have to specify which template to load
	if err != nil {
		log.Fatalln(err)
	}
}