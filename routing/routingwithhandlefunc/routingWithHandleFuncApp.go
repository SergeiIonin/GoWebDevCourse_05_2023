package main

import (
	"io"
	"net/http"
)

func helloRoute(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello!!!")
}

func whatsupRoute(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "whatsup???")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloRoute)
	mux.HandleFunc("/whatsup", whatsupRoute)

	http.ListenAndServe(":8080", mux)
}
