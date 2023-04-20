package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln("Failed to parse file", err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Failed to create a file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)

	if err != nil {
		log.Fatalln(err)
	}

}
