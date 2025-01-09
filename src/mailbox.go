package main

import (
	"net/http"
)

type MailboxTmplData struct {
	Letters []Letter
	User User
}

func handleMailbox(w http.ResponseWriter, r *http.Request) {
	data := MailboxTmplData{}
	data.User = sessionUser(w, r)
	data.Letters = carriedLettersFromCache(w, data.User.MailboxCache)
	renderTemplate(w, "mailbox", data)
}

func handleSent(w http.ResponseWriter, r *http.Request) {
	data := MailboxTmplData{}
	data.User = sessionUser(w, r)
	data.Letters = lettersFromCache(w, data.User.SentCache)
	renderTemplate(w, "sent", data)
}
