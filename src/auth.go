package main

import (
	"os"
	"net/http"
	"github.com/google/uuid"
)

const TokenDir = "../../data/tokens/"
const TokenFileExt = ".tk"
const TokenCookieStr = "sessiontoken"
const HandleCookieStr = "sessionhandle"

func handlePostSignup(w http.ResponseWriter, r *http.Request) {
	// TODO: Check for duplicate users (dupl. emails fine)
	u := &User{
		Handle: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email: r.FormValue("email"),
		// Checkboxes aren't sent as form values if they aren't checked, hence:
		NotifyByEmail: r.FormValue("notifybyemail") != ""}
	u.save()
	http.Redirect(w, r, "/confirmsignup/", http.StatusFound)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	u, err := loadUser(r.FormValue("username"))
	if err != nil || u.Password != r.FormValue("password") {
		http.Redirect(w, r, "/landing/badlogin/", http.StatusFound)
		return
	}

	tokenCookie := &http.Cookie{
        Name:  TokenCookieStr,
        Value: newToken(u.Handle),
        Path:  "/"}
    http.SetCookie(w, tokenCookie)

	handleCookie := &http.Cookie{
        Name:  HandleCookieStr,
        Value: u.Handle,
        Path:  "/"}
    http.SetCookie(w, handleCookie)

	http.Redirect(w, r, "/mailbox/", http.StatusFound)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	err := clearSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/landing/", http.StatusFound)
}

func newToken(handle string) string {
	id := uuid.NewString()
	fname := TokenDir + handle + TokenFileExt
	os.WriteFile(fname, []byte(id), 0600)
	return id
}

func clearSession(r *http.Request) error {
	handle, err := storedHandle(r)
	if err != nil {
		return err
	}

	err = os.Remove(TokenDir + handle + TokenFileExt)
	if err != nil {
		return err
	}
	return nil
}

func sessionUser(w http.ResponseWriter, r *http.Request) User {
	u := User{}

	h, err := storedHandle(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if !storedTokenValid(r, h) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	u, err = loadUser(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return u
}

func storedHandle(r *http.Request) (string, error) {
	handle, err := r.Cookie(HandleCookieStr)
	if err != nil {
		return "", err
	}
	return handle.Value, nil
}

func storedTokenValid(r *http.Request, handle string) bool {
	id, err := r.Cookie(TokenCookieStr)
	if err != nil {
		return false
	}

	fname := TokenDir + handle + TokenFileExt
	storedId, err := os.ReadFile(fname)
	if err != nil {
		return false
	}

	if id.Value != string(storedId) {
		// TODO: delete cookie if not matched?
		return false
	}
	return true
}
