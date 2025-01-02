package main

import "net/http"

func handleLanding(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "landing")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "login")
}
