package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//var sessions map[string]int

const visitsNumberCookie = "VisitsNumber"

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/visits", visits)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(visitsNumberCookie)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  visitsNumberCookie,
			Value: "0",
			Path:  "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln("Can't get the count for user: ", err.Error()) // todo how to recover from this?
	} else {
		count++
		fmt.Println("the current count is ", count)
		cookie.Value = strconv.Itoa(count)

		http.SetCookie(w, cookie)
	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func visits(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(visitsNumberCookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
