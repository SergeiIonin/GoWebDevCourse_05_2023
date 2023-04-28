package main

import (
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// run it by
// go run *.go
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	// logged in?
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return // NB
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")

		// if username is already taken, then return 403
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username is already taken!", http.StatusForbidden)
			return
		}

		// set cookie
		sId := uuid.New().String()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sId,
		}
		http.SetCookie(w, cookie)

		// save session
		dbSessions[sId] = un

		// save user
		u := user{un, pwd, fn, ln}
		dbUsers[un] = u

		// redirect when all is done
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return // NB
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
