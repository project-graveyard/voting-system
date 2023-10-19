package main

import (
	"log"
	"net/http"

	"github.com/DaveSaah/voting-system/handlers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/login", handlers.Login)

	log.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
