package main

import "net/http"

type LandingTemplateData struct {
	Users []User
	BadLogin bool
}

// Landing page (pre-sign in)
func handleLanding(w http.ResponseWriter, r *http.Request) {
	handle, err := storedHandle(r)
	if err == nil {
		if storedTokenValid(r, handle) {
			http.Redirect(w, r, "/mailbox/" + handle, http.StatusFound)
		}
	}

	tmplData := LandingTemplateData{BadLogin: false}	

	suffix := r.URL.Path[len("/landing/"):]
	if suffix == "badlogin/" {
		tmplData.BadLogin = true	
	}

	tmplData.Users, err = getUserList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = templates.ExecuteTemplate(w, "landing.html", tmplData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
