package main

import "net/http"

type LandingTemplateData struct {
	Users []User
	BadLogin bool
}

// Landing page (pre-sign in)
func handleLanding(w http.ResponseWriter, r *http.Request) {
	tmplData := LandingTemplateData{BadLogin: false}	

	suffix := r.URL.Path[len("/landing/"):]
	if suffix == "badlogin/" {
		tmplData.BadLogin = true	
	}

	var err error
	tmplData.Users, err = getUserList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = templates.ExecuteTemplate(w, "landing.html", tmplData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
