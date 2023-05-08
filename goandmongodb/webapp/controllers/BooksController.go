package controllers

import (
	"GoWebDevCourse/goandmongodb/webapp/books"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type BooksController struct {
	template *template.Template
}

func NewBooksController(tpl *template.Template) *BooksController {
	return &BooksController{
		template: tpl,
	}
}

func (c BooksController) BooksIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	booksFetched, err := books.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500)+"Error executing request "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	c.template.ExecuteTemplate(w, "books.gohtml", booksFetched)
}

func (c BooksController) BooksShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookFetched, err := books.OneBook(r)
	if err != nil {
		http.Error(w, "Error executing request "+err.Error(), http.StatusInternalServerError) // todo it's not always 500!
		return
	}

	w.WriteHeader(http.StatusOK)
	c.template.ExecuteTemplate(w, "show.gohtml", bookFetched)
}

func (c BooksController) BooksCreateForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c.template.ExecuteTemplate(w, "create.gohtml", nil)
}

func (c BooksController) CreateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreated, err := books.PutBook(r)
	if err != nil {
		http.Error(w, "Error executing request "+err.Error(), http.StatusInternalServerError) // todo it's not always 500!
		return
	}

	c.template.ExecuteTemplate(w, "created.gohtml", bookCreated)
}

func (c BooksController) UpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreated, err := books.OneBook(r)
	if err != nil {
		http.Error(w, "Error executing request "+err.Error(), http.StatusInternalServerError) // todo it's not always 500!
		return
	}

	c.template.ExecuteTemplate(w, "update.gohtml", bookCreated)
}

func (c BooksController) UpdateBookProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreated, err := books.UpdateBook(r)
	if err != nil {
		http.Error(w, "Error executing request "+err.Error(), http.StatusInternalServerError) // todo it's not always 500!
		return
	}

	c.template.ExecuteTemplate(w, "updated.gohtml", bookCreated)
}

func (c BooksController) DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := books.DeleteBook(r)
	if err != nil {
		http.Error(w, "Error executing request "+err.Error(), http.StatusInternalServerError) // todo it's not always 500!
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
