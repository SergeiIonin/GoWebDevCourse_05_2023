package main

import (
	"GoWebDevCourse/goandmongodb/webapp/controllers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

import "html/template"

var tpl *template.Template
var db *mongo.Database
var booksCollection *mongo.Collection

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := httprouter.New()
	bc := controllers.NewBooksController(tpl)

	r.GET("/", index)
	r.GET("/books", bc.BooksIndex)
	r.GET("/favicon.ico", notFound)

	r.GET("/books/show", bc.BooksShow)
	r.GET("/books/create", bc.BooksCreateForm)
	r.POST("/books/create/process", bc.CreateBook)
	r.POST("/books/update", bc.UpdateBook)
	r.POST("/books/update/process", bc.UpdateBookProcess)
	r.DELETE("/books/delete/process", bc.DeleteBook)

	http.ListenAndServe(":8080", r)
}

func notFound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
