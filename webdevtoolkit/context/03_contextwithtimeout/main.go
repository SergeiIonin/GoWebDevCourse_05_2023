package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel() // NB cancel is a func

	ch := make(chan int)

	go func() {
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second) // by making this value more or less than timeout we're changing the output (will be either 200 or 408)

		// return immediately if there's an error
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}() // NB we need to actually call this anonymous func by adding ()

	// either the ctx is Done and we return 408 or we read from the channel and return 200
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}
