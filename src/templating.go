package main

import(
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func executeStaticTemplate(w http.ResponseWriter, fname string) {
	err := templates.ExecuteTemplate(w, fname, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
