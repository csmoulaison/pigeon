package main

import(
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
