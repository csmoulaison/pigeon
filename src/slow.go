package main

import (
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/landing/", http.StatusFound)
}

func main() {
	// TODO: Can we abstract these locked handlers further, indirecting the call to
	// http.HandleFunc(), for instance?
	http.HandleFunc("/",                handleIndex)
	http.HandleFunc("/landing/",        handleLanding)
	http.HandleFunc("/signup/",         staticHandler("signup"))
	http.HandleFunc("/confirmsignup/",  staticHandler("confirmsignup"))
	http.HandleFunc("/postsignup/",     handlePostSignup)
	http.HandleFunc("/login/",          handleLogin)
	http.HandleFunc("/logout/",         handleLogout)

	http.HandleFunc("/mailbox/",       	lockedHandler(handleMailbox, "/mailbox/"))
	http.HandleFunc("/deletemail/",     lockedHandler(handleDeleteMail, "/deleteemail/"))
	http.HandleFunc("/sent/",           lockedHandler(handleSent, "/sent/"))
	http.HandleFunc("/deletesent/",     lockedHandler(handleSent, "/deletesent/"))
	http.HandleFunc("/rolodex/",        lockedHandler(handleSent, "/rolodex/"))
	http.HandleFunc("/deletecontact/",  lockedHandler(handleSent, "/deletecontact/"))
	http.HandleFunc("/settings/",       lockedHandler(handleSent, "/settings/"))
	http.HandleFunc("/modifysettings/", lockedHandler(handleSent, "/modifysettings/"))
	http.HandleFunc("/send/",           lockedHandler(handleSent, "/send/"))
	http.HandleFunc("/postsend/",       lockedHandler(handleSent, "/postsend/"))
	http.HandleFunc("/confirmsend/",    lockedHandler(handleSent, "/confirmsend/"))

	// TODO: Maybe command line arg for port?
	log.Fatal(http.ListenAndServe(":8080", nil))
}
