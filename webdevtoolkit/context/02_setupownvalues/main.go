package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// we're enriching the ctx using parent context
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dbAccess(ctx)

	fmt.Fprintln(w, results)
}

// it'll return smth like this
// context.Background.WithValue(type *http.contextKey, val <not Stringer>).WithValue(type *http.contextKey, val 127.0.0.1:8080).WithCancel.WithCancel
func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// accessing the context value by key and verifying that the type is int
func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userID").(int) // NB!
	return uid
}
