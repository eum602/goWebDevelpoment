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
	err := tpl.ExecuteTemplate(os.Stdout,"tpl1.goHtml","Erick") //here we can pass also aggregate data like slices, structs and so on
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout,"tpl2.goHtml","eum602") //with ExecuteTemplate we have to specify which template to load
	if err != nil {
		log.Fatalln(err)
	}
}