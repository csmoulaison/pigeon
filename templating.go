package main

import(
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles("landing.html", "login.html"))

func executeTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
