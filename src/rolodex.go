package main

import (
	"net/http"
)

func handleRolodex(w http.ResponseWriter, r *http.Request) {
	u := sessionUser(w, r)
	renderTemplate(w, "rolodex", u)
}

func handleAddContact(w http.ResponseWriter, r *http.Request) {
	u := sessionUser(w, r)
	h := r.FormValue("handle")
	u.Rolodex = append(u.Rolodex, h)
	u.save()	
	http.Redirect(w, r, "/rolodex/", http.StatusFound)
}

func handleDeleteContact(w http.ResponseWriter, r *http.Request) {
	u := sessionUser(w, r)
	h := r.FormValue("handle")
	for i, contact := range u.Rolodex {
        if contact == h {
            u.Rolodex = append(u.Rolodex[:i], u.Rolodex[i+1:]...)
            break
        }
    }
	u.save()	
	http.Redirect(w, r, "/rolodex/", http.StatusFound)
}
