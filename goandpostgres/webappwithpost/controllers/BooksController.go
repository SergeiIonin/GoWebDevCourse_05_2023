package controllers

import (
	"GoWebDevCourse/goandpostgres/webappwithget/model"
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type BooksController struct {
	session  *sql.DB
	template *template.Template
}

func NewBooksController(source string, tpl *template.Template) *BooksController {
	// "postgres://sergei:123@localhost/bookstore?sslmode=disable"
	db, err := sql.Open("postgres", source)
	if err != nil {
		panic(err)
	}

	//defer db.Close() // todo when do we need to close?

	return &BooksController{
		session:  db,
		template: tpl,
	}
}

func (c BooksController) BooksIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rows, err := c.session.Query("SELECT * FROM books;")
	if err != nil {
		http.Error(w, "Error executing request"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	books := make([]model.Book, 0)
	for rows.Next() {
		var bk model.Book
		if err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price); err != nil {
			panic(err)
		}
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	c.template.ExecuteTemplate(w, "books.gohtml", books)
}

func (c BooksController) BooksShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, "Wrong path", http.StatusBadRequest)
		return
	}

	row := c.session.QueryRow("SELECT * FROM books WHERE isbn=$1;", isbn)

	var book model.Book

	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	if err != nil {
		http.Error(w, fmt.Sprintf("Book with ISBN %s", isbn), http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)

	c.template.ExecuteTemplate(w, "show.gohtml", book)
}

func (c BooksController) BooksCreateForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c.template.ExecuteTemplate(w, "create.gohtml", nil)
}

func (c BooksController) CreateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get form values
	bk := model.Book{}

	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	price := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || price == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable) // NB 406
		return
	}
	bk.Price = float32(f64)

	// insert values
	_, err = c.session.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	c.template.ExecuteTemplate(w, "created.gohtml", bk)
}
