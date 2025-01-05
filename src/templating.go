package main

import(
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
