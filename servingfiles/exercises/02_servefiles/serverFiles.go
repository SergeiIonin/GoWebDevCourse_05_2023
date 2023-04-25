package main

import (
	"log"
	"net/http"
)

func main() {
	//This way, the index.html will be displayed
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./starting-files"))))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
