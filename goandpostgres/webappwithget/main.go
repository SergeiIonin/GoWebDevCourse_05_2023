package main

import (
	"GoWebDevCourse/goandpostgres/webappwithget/controllers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	source := "postgres://sergei:123@localhost/bookstore?sslmode=disable"
	r := httprouter.New()
	bc := controllers.NewBooksController(source)
	r.GET("/books", bc.GetBooks)
	r.GET("/books/:isbn", bc.GetBookByISBN)

	http.ListenAndServe(":8080", r)

}
