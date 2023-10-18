package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Get" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("%s tried to access", r.RemoteAddr)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		log.Println("Form data received")
		// email := r.FormValue("email")
		// password := r.FormValue("password")
		return
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to the dashboard")
	}
}
