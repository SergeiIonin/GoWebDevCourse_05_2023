package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func main() {

	r := httprouter.New()
	r.GET("/items", items)
	r.DELETE("/items/delete", delete)

	http.ListenAndServe(":8080", r)

}

func items(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	qp := req.URL.Query().Get("id")
	io.WriteString(w, "Your items will be here, you requested id "+qp)
}

func delete(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	io.WriteString(w, "Your items will be deleted, ha-ha-ha")
}
