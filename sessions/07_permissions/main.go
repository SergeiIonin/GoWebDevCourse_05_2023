package main

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "Only 007 can access that endpoint!", http.StatusForbidden)
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

	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		r := req.FormValue("role")

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
		pEnc, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		u = user{un, pEnc, fn, ln, r}
		dbUsers[un] = u

		// redirect when all is done
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return // NB
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	// logged in?
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password don't match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pwd))
		if err != nil {
			http.Error(w, "password is wrong", http.StatusForbidden)
			return
		}

		sId := uuid.New().String()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sId,
		}
		http.SetCookie(w, cookie)
		dbSessions[sId] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	cookie, err := req.Cookie("session")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError) // this is unlikely?
		return
	}

	// NB, set the value to blank
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	delete(dbSessions, cookie.Value) // NB
	http.SetCookie(w, cookie)

	http.Redirect(w, req, "/login", http.StatusSeeOther) // NB redirect to /login
}
