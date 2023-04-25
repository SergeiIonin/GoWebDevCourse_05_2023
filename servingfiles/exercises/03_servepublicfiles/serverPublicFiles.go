package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func pics(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "Error executing template", 500)
	}
}

func main() {
	http.HandleFunc("/", pics)
	http.Handle("/pics/", http.FileServer(http.Dir("./starting-files/public"))) // NOT just /pics !
	http.ListenAndServe(":8080", nil)
}
