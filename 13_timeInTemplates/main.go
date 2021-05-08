package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.goHtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
	//01/02 03:04:05 MST 2006 -> january (01 is month) second (02 is the day), 03 hours (03 is hours) 04 minutes (04 are the minutes) 05 seconds (05 are the seconds) of year 2006 (2006 represents the year)

	//for example if I put t.Format("02-01-2006") then if means the day goes first the the month and then the year
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"fdateDMY": dayMonthYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.goHtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	//RUN WITH: go run main.go > index.html &&  chromium-browser index.html && rm index.html
}
