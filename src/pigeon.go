package main

// TODO:
// - email
// - delay

import (
	"log"
	"net/http"
	"io/ioutil"

	//"golang.org/x/crypto/acme/autocert"
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

func ttfRoute(p string, i string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile("assets/" + i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "font/ttf")
		w.Write(buf)
	}
	http.HandleFunc("/" + p + "/", handler)
}

func main() {
	// Pre login
	http.HandleFunc("/", handleIndex)
	route("landing", handleLanding)
	staticRoute("signup")
	staticRoute("confirmSignup")
	route("postsignup", handlePostSignup)
	route("login", handleLogin)
	route("logout", handleLogout)
	// Post login
	lockedRoute("mailbox", handleMailbox)
	lockedRoute("sent", handleSent)
	lockedRoute("view", handleView)
	lockedRoute("rolodex", handleRolodex)
	lockedRoute("addcontact", handleAddContact)
	lockedRoute("deletecontact", handleDeleteContact)
	lockedRoute("settings", handleSettings)
	lockedRoute("postsettings", handlePostSettings)
	lockedRoute("send", handleSend)
	lockedRoute("postsend", handlePostSend)
	lockedRoute("confirmsend", handleConfirmSend)

	pngRoute("pigeon", "pigeon.png")
	ttfRoute("proggy", "ProggyVector-Regular.ttf")

	cert := &autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("hellopigeon.net", "www.hellopigeon.net"),
	}
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: cert.TLSConfig(),
	}

	log.Fatal(s.ListenAndServeTLS("", ""))

	//log.Fatal(http.ListenAndServe(":80", nil))
}
