package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"getMonthDayYear": getMonthDayYear,
}

func getMonthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
