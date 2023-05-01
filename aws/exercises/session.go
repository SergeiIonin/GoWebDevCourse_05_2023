package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLen // NB outside of if err != nil {...}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if us, ok := dbSessions[c.Value]; ok {
		us.lastActivity = time.Now()
		dbSessions[c.Value] = us
		u = dbUsers[us.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	us, ok := dbSessions[c.Value]
	if ok {
		// update lastActivity in the user session
		us.lastActivity = time.Now()
		dbSessions[c.Value] = us
	}
	// refresh session
	http.SetCookie(w, c)
	c.MaxAge = sessionLen
	_, ok = dbUsers[us.un]
	return ok
}

func cleanSessions() {
	fmt.Println("Sessions before clean:")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}

	dbSessionCleaned = time.Now()
	fmt.Println("Sessions after clean:")
	showSessions()
}

func showSessions() {
	fmt.Println("*****")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
