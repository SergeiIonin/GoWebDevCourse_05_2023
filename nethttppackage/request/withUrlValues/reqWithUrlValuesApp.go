package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
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

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}

	tpl.Execute(w, data)
}

func main() {
	var handler hdlr
	http.ListenAndServe(":8080", handler)
}
