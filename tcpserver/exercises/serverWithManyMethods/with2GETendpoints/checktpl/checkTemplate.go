package main

import (
	"log"
	"os"
	"text/template"
)

type response struct {
	Title string
	H1    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	resp := response{
		Title: "Hello",
		H1:    "HEY!!!",
	}
	err := tpl.Execute(os.Stdout, resp)
	if err != nil {
		log.Fatalln(err)
	}

	f, _ := os.Create("file")
	err1 := tpl.Execute(f, resp)
	if err1 != nil {
		log.Fatalln("error exe tpl, ", err1)
	}
}
