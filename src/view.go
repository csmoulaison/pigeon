package main

import (
	"net/http"
	"strconv"
	"bufio"
	"strings"
)

type ViewTmplData struct {
	Letter Letter
	Paragraphs []string
	User User
}

func handleView(w http.ResponseWriter, r *http.Request) {
	u := sessionUser(w, r)

	q := r.URL.Query()
	idStr := q.Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		renderTemplate(w, "badview", u)
		return
	}

	l := Letter{}
	letter_matched := false

	for _, cached := range u.MailboxCache {
		if cached == id {
			l, err = loadLetter(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			letter_matched = true
			break
		}
	}

	for _, cached := range u.SentCache {
		if cached == id {
			l, err = loadLetter(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			letter_matched = true
			break
		}
	}

	if !letter_matched {
		renderTemplate(w, "badview", u)
		return
	}

	paras := []string{}
	s := bufio.NewScanner(strings.NewReader(l.Body))
	for s.Scan() {
		paras = append(paras, s.Text())
	}

	l.Read = true
	l.save()

	data := ViewTmplData{Letter: l, User: u, Paragraphs: paras}
	renderTemplate(w, "view", data)
}
