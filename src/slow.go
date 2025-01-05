package main

import (
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/landing/", http.StatusFound)
}

// Two things we want to accomplish right now:
// 1. Have a way to easily configure a route to pass session handle to a route
// 2. Abstract http.HandleFunc further to reduce redundancy

func route(p string, f http.HandlerFunc) {
	http.HandleFunc("/" + p + "/", f)
}

func staticRoute(p string) {
	f := func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, p, nil)
	}
	http.HandleFunc("/" + p + "/", f)
}

func lockedRoute(p string, f http.HandlerFunc) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		handle := r.URL.Path[len("/" + p + "/"):]
		if !storedTokenValid(r, handle) {
			http.Redirect(w, r, "/landing/", http.StatusFound)
		}
		f(w, r)
	}
	http.HandleFunc("/" + p + "/", handler)
}

func main() {
	route("/", handleIndex)
	route("landing", handleLanding)
	staticRoute("signup")
	staticRoute("confirmSignup")
	route("postsignup", handlePostSignup)
	route("login", handleLogin)
	route("logout", handleLogout)

	lockedRoute("mailbox", handleMailbox)
	lockedRoute("deletemail", handleDeleteMail)
	lockedRoute("sent", handleSent)
	lockedRoute("deletesent", handleSent)
	lockedRoute("rolodex", handleSent)
	lockedRoute("deletecontact", handleSent)
	lockedRoute("settings", handleSent)
	lockedRoute("modifysettings", handleSent)
	lockedRoute("send", handleSent)
	lockedRoute("postsend", handleSent)
	lockedRoute("confirmsend", handleSent)

	// TODO: Maybe command line arg for port?
	log.Fatal(http.ListenAndServe(":8080", nil))
}
