package main

import (
	"io"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	// this endpoint is called when we reference it in html:
	//<img src="/dog.jpg" alt="a picture of a dog">
	http.HandleFunc("/dog.jpg", chien)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Content-Type", "text/html; charset=utf-8")
	resp := "<h1>foo ran</h1>"
	io.WriteString(w, resp)
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		http.Error(w, "Not Found", 404)
		return
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
