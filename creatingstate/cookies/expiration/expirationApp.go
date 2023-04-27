package main

import (
	"fmt"
	"net/http"
)

const sessionCookie = "session"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:  sessionCookie,
		Value: "session is valid",
	}
	http.SetCookie(w, c)
	fmt.Fprintln(w, `<h1><a href="/read">read a cookie</h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(sessionCookie)
	if err == http.ErrNoCookie {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your cookie is <br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(sessionCookie)
	if err == http.ErrNoCookie {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
