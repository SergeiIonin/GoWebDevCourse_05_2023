package controllers

import (
	"GoWebDevCourse/goandpostgres/webappwithget/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type BooksController struct {
	session *sql.DB
}

func NewBooksController(source string) *BooksController {
	// "postgres://sergei:123@localhost/bookstore?sslmode=disable"
	db, err := sql.Open("postgres", source)
	if err != nil {
		panic(err)
	}

	//defer db.Close() // todo when do we need to close?

	return &BooksController{
		session: db,
	}
}

func (c BooksController) GetBooks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	payload, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(payload))
	//w.Write(payload)
}

func (c BooksController) GetBookByISBN(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	isbn := p.ByName("isbn")
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

	payload, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//io.WriteString(w, string(payload))
	w.Write(payload)
}
