package main

import (
	"GoWebDevCourse/goandpostgres/webappwithpost/controllers"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var db *sql.DB
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// export fields to templates
// fields changed to uppercase
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func main() {
	source := "postgres://sergei:123@localhost/bookstore?sslmode=disable"
	r := httprouter.New()
	bc := controllers.NewBooksController(source, tpl)

	r.GET("/", index)
	r.GET("/books", bc.BooksIndex)
	r.GET("/books/show", bc.BooksShow)
	r.GET("/books/create", bc.BooksCreateForm)
	r.POST("/books/create/process", bc.CreateBook)

	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
