package main

import "net/http"

// Auth handlers
func handleSignup(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleConfirmSignup(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "confirmsignup.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlePostSignup(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Handle: r.FormValue("username"),
		Password: r.FormValue("password"),
		DisplayName: r.FormValue("displayname"),
		Email: r.FormValue("email"),
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
