package main

import (
	"net/http"
)

type MailboxTmplData struct {
	Letters []Letter
	User User
}

// Mailbox (received)
func handleMailbox(w http.ResponseWriter, r *http.Request) {
	data := MailboxTmplData{}
	data.User = sessionUser(w, r)
	data.Letters = lettersFromCache(w, data.User.MailboxCache)
	renderTemplate(w, "mailbox", data)
}

func handleDeleteMail(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Sent
func handleSent(w http.ResponseWriter, r *http.Request) {
	data := MailboxTmplData{}
	data.User = sessionUser(w, r)
	data.Letters = lettersFromCache(w, data.User.SentCache)
	renderTemplate(w, "sent", data)
}

func handleDeleteSent(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

