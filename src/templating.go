package main

import(
	"html/template"
)

const TemplateDir = "./templates/"

var templates = template.Must(template.ParseFiles(
	TemplateDir + "landing.html", 
	TemplateDir + "login.html",
	TemplateDir + "mailbox.html"))
