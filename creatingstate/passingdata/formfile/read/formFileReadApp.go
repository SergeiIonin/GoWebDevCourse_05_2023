package main

import (
	"fmt"
	"io"
	"net/http"
)

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
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	<br>`+s)
}
