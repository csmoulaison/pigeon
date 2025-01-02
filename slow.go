package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/landing/", handleLanding)
	http.HandleFunc("/login/", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
