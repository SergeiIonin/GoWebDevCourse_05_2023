package main

import (
	"io"
	"net/http"
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
	req.Header.Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>Sergei</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler) // we are registering a new handler on the DefaultServerMux
	http.HandleFunc("/me", meHandler)
	http.HandleFunc("/dog", dogHandler)

	http.ListenAndServe(":8080", nil)
}
