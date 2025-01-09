package main

import "net/http"

type SettingsTmplData struct {
	User User
	JustSaved bool
}

func handleSettings(w http.ResponseWriter, r *http.Request) {
	data := SettingsTmplData{}
	data.User = sessionUser(w, r)

	suffix := r.URL.Path[len("/settings/"):]
	data.JustSaved = suffix == "saved/"

	renderTemplate(w, "settings", data)
}

func handlePostSettings(w http.ResponseWriter, r *http.Request) {
	u := sessionUser(w, r)

	p := r.FormValue("password")
	if p != "" {
		u.Password = p
	}

	u.Email = r.FormValue("email")
	u.NotifyByEmail = r.FormValue("notifybyemail") != ""

	u.save()
	http.Redirect(w, r, "/settings/saved/", http.StatusFound)
}
