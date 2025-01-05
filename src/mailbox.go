package main

import (
	"net/http"
)

const MailboxRoute = "/mailbox/"

func handleMailbox(w http.ResponseWriter, r *http.Request) {
	handle := r.URL.Path[len(MailboxRoute):]
	err := templates.ExecuteTemplate(w, "mailbox.html", handle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDeleteMail(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}
