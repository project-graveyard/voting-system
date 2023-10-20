package handlers

import (
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("%s tried to access", r.RemoteAddr)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		log.Println("Form data received")
		return
	}
}
