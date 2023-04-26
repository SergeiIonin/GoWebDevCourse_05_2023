package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func processError(w http.ResponseWriter, err error, prefix string) {
	if err != nil {
		fmt.Println(prefix + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)

	if req.Method == http.MethodPost {

		file, header, err := req.FormFile("q")
		processError(w, err, "error calling FormFile: ")

		defer file.Close()

		fmt.Println("\nfile:", file, "\nheader:", header, "\nerr", err)

		bs, err := io.ReadAll(file)
		processError(w, err, "Error reading file: ")

		s = string(bs)

		dst, err := os.Create(filepath.Join("./user/", header.Filename))
		processError(w, err, "error creating the dst file: ")

		defer dst.Close()

		_, err = dst.WriteString(s)
		processError(w, err, "error writing to the dst file: ")

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}
