package main

import (
	"github.com/google/uuid"
	"net/http"
	"text/template"
)

var tpl *template.Template

type user struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = make(map[string]user)
var dbSessions = make(map[string]string)

const sessionCookie = "Session"

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users", users)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(sessionCookie)
	if err != nil {
		sID := uuid.New()
		cookie = &http.Cookie{
			Name:  sessionCookie,
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		first := req.FormValue("firstname")
		last := req.FormValue("lastname")
		u := user{username, first, last}
		dbSessions[cookie.Value] = username
		dbUsers[username] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func users(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(sessionCookie)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	un, ok := dbSessions[cookie.Value]

	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "user.gohtml", u)
}
