package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func foo(w http.ResponseWriter, req *http.Request) {
	firstName := req.FormValue("first")
	lastName := req.FormValue("last")
	subscribed := req.FormValue("subscribe") == "on"

	person := person{firstName, lastName, subscribed}

	err := tpl.ExecuteTemplate(w, "index.gohtml", person)

	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

}
