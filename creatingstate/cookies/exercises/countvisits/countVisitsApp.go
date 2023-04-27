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
	c, err := req.Cookie(visitsNumberCookie)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  visitsNumberCookie,
			Value: "1",
			Path:  "/",
		})
	} else {
		count, err := strconv.Atoi(c.Value)
		if err != nil {
			log.Fatalln("Can't get the count for user: ", err.Error()) // todo how to recover from this?
		} else {
			count++
			fmt.Println("the current count is ", count)
			http.SetCookie(w, &http.Cookie{
				Name:  visitsNumberCookie,
				Value: fmt.Sprint(count),
				Path:  "/",
			})
		}
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
