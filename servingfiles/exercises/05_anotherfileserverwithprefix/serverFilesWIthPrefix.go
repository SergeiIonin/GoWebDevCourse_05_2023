package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "Error executing template", 500)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.FileServer(http.Dir("./starting-files")))
	http.ListenAndServe(":8080", nil)
}
