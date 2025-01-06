package main

import (
	"net/http"
	"time"
)

func handleSend(w http.ResponseWriter, r *http.Request) {
	u := sessionUser(w, r)
	renderTemplate(w, "send", u)
}

func handlePostSend(w http.ResponseWriter, r *http.Request) {
	// Store letter with id
	l := Letter{}
	var err error
	l.Id, err = newLetterId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	l.Title = r.FormValue("title")
	l.Body = r.FormValue("body")
	l.Created = time.Now()

	// Cache letter id with sender
	sender := sessionUser(w, r)
	sender.SentCache = append(sender.SentCache, l.Id)
	err = sender.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Cache letter id with recipient
	h := r.FormValue("recipient")
	recipient := User{}
	recipient, err = loadUser(h)
	// TODO: Should we tell the user if this recipient doesn't exist?
	// As of now, we are just doing nothing if the recipient doesn't exist, and the
	// letter still appears in the users sent box.
	if err == nil {
		recipient.MailboxCache = append(recipient.MailboxCache, l.Id)
		err = recipient.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// Save letter after setting sender and recipient
	l.Sender = sender.Handle
	l.Recipient = recipient.Handle
	err = l.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/confirmsend/", http.StatusFound)
}

func handleConfirmSend(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "confirmsend", nil)
}
