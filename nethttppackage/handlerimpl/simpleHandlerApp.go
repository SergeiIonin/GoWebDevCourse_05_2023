package main

import (
	"fmt"
	"net/http"
)

type hdlr string

func (m hdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HEY!!!")
}

func main() {
	var handler hdlr
	http.ListenAndServe(":8090", handler)

}
