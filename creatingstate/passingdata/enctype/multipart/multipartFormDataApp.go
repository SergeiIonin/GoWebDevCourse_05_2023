package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)

	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
	// the output will look like
	/*BODY: -----------------------------14145121929322018272566465981 Content-Disposition: form-data; name="first"
	Tom -----------------------------14145121929322018272566465981 Content-Disposition: form-data; name="last"
	Moore -----------------------------14145121929322018272566465981 Content-Disposition: form-data;
	name="subscribe" on -----------------------------14145121929322018272566465981--*/
}
