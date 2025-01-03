package main

import (
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/landing/", http.StatusFound)
}

func main() {
	http.HandleFunc("/",                handleIndex)
	http.HandleFunc("/landing/",        handleLanding)
	http.HandleFunc("/signup/",         handleSignup)
	http.HandleFunc("/postsignup/",     handlePostSignup)
	http.HandleFunc("/login/",          handleLogin)
	http.HandleFunc("/logout/",         handleLogout)
	http.HandleFunc("/mailbox/",        handleMailbox)
	http.HandleFunc("/deletemail/",     handleDeleteMail)
	http.HandleFunc("/sent/",           handleSent)
	http.HandleFunc("/deletesent/",     handleDeleteSent)
	http.HandleFunc("/rolodex/",        handleRolodex)
	http.HandleFunc("/deletecontact/",  handleDeleteContact)
	http.HandleFunc("/settings/",       handleDeleteMail)
	http.HandleFunc("/modifysettings/", handleModifySettings)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
