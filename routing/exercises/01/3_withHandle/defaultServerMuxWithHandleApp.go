package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>Hello, this is the main page</h1>")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>DOG</h1>")
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, req.Form["name"][0])
}

func mev2Handler(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "hello.gohtml", "Sergei")
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("hello.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(indexHandler)) // we are registering a new handler on the DefaultServerMux
	http.Handle("/me", http.HandlerFunc(meHandler))
	http.Handle("/v2/me", http.HandlerFunc(mev2Handler))
	http.Handle("/dog", http.HandlerFunc(dogHandler))

	http.ListenAndServe(":8080", nil)
}
