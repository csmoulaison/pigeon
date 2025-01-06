package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/landing/", http.StatusFound)
}

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
		handle, err := storedHandle(r)
		if err != nil || !storedTokenValid(r, handle) {
			http.Redirect(w, r, "/landing/", http.StatusFound)
		}
		f(w, r)
	}
	http.HandleFunc("/" + p + "/", handler)
}

func pngRoute(p string, i string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile("assets/" + i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)
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
	lockedRoute("deletesent", handleDeleteSent)
	lockedRoute("rolodex", handleRolodex)
	lockedRoute("deletecontact", handleDeleteContact)
	lockedRoute("settings", handleSettings)
	lockedRoute("postsettings", handlePostSettings)
	lockedRoute("send", handleSend)
	lockedRoute("postsend", handlePostSend)
	lockedRoute("confirmsend", handleConfirmSend)

	pngRoute("pigeon", "pigeon.png")

	// TODO: Maybe command line arg for port?
	log.Fatal(http.ListenAndServe(":8080", nil))
}
