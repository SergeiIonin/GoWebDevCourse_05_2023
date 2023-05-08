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
	r.GET("/books/update", bc.UpdateBook)
	r.POST("/books/update/process", bc.UpdateBookProcess)
	r.GET("/books/delete", deleteHandle)
	// we can use DELETE if we will specify the method from the JS script call,
	// here's the Chat GPT proposal (NOT TESTED)
	// <a href="#" onclick="deleteResource('/resource/123')">Delete</a>
	//
	//<script>
	//function deleteResource(url) {
	//  const request = new XMLHttpRequest();
	//  request.open('DELETE', url);
	//  request.send();
	//}
	//</script>
	r.GET("/books/delete/process", bc.DeleteBook) // fixme if the method is DELETE, I get the message 405 Method Not Allowed

	http.ListenAndServe(":8080", r)
}

func notFound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func deleteHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/books/delete/process", http.StatusSeeOther)
}
