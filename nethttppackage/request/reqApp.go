package main

import (
	"html/template"
	"log"
	"net/http"
)

type hdlr string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (h hdlr) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(w, req.Form)
}

func main() {
	var handler hdlr
	http.ListenAndServe(":8080", handler)
}
