package main

import (
	"io"
	"net/http"
)

func main() {
	//http.StripPrefix returns a handler that will process requests with the path starting with prefix, the dir
	// for the FileServer could be arbitrary
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// NOT /assets/toby.jpg
	io.WriteString(w, `
	<img src="/resources/toby.jpg">
	`)
}
