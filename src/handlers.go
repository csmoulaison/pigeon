package main

import "net/http"

// TODO: Split and deprecate this file per concern. The current split by comment
// is probably fine.
// - sent.go
// - rolodex.go
// - settings.go
// - send.go

func staticHandler(tmpl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, tmpl, nil)
	}
}

func lockedHandler(fn http.HandlerFunc, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle := r.URL.Path[len(path):]
		if !storedTokenValid(r, handle) {
			http.Redirect(w, r, "/landing/", http.StatusFound)
		}
		fn(w, r)
	}
}

// Sent
func handleSent(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleDeleteSent(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Rolodex
func handleRolodex(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleDeleteContact(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Settings
func handleSettings(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleModifySettings(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Send mail
func handleSend(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handlePostSend(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleConfirmSend(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}
