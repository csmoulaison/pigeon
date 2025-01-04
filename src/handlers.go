package main

import "net/http"

// TODO: Split and deprecate this file per concern. The current split by comment
// is probably fine.
// - auth.go
// - mailbox.go
// - sent.go
// - rolodex.go
// - settings.go
// - send.go

// Auth handlers
func handleSignup(w http.ResponseWriter, r *http.Request, ) {
	executeStaticTemplate(w, "signup.html")
}

func handleConfirmSignup(w http.ResponseWriter, r *http.Request) {
	executeStaticTemplate(w, "confirmsignup.html")
}

func handlePostSignup(w http.ResponseWriter, r *http.Request) {
	// TODO: Check for duplicate users (dupl. emails fine)
	u := &User{
		Handle: r.FormValue("username"),
		Password: r.FormValue("password"),
		DisplayName: r.FormValue("displayname"),
		Email: r.FormValue("email"),
		// Checkboxes aren't sent as form values if they aren't checked, hence:
		NotifyByEmail: (r.FormValue("notifybyemail") != "")}
	u.save()
	http.Redirect(w, r, "/confirmsignup/", http.StatusFound)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	u, err := loadUser(r.FormValue("username"))
	if err != nil || u.Password != r.FormValue("password") {
		http.Redirect(w, r, "/landing/badlogin/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/mailbox/" + u.Handle, http.StatusFound)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: probably clear browser cached info, or something?
	http.Redirect(w, r, "/landing/", http.StatusOK)
}

// Mailbox
func handleMailbox(w http.ResponseWriter, r *http.Request) {
	handle := r.URL.Path[len("/mailbox/"):]
	err := templates.ExecuteTemplate(w, "mailbox.html", handle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDeleteMail(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Sent
func handleSent(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleDeleteSent(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Rolodex
func handleRolodex(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleDeleteContact(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Settings
func handleSettings(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleModifySettings(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// Send mail
func handleSend(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handlePostSend(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func handleConfirmSend(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}
