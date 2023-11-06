package main

import (
	"log"
	"net/http"

	"github.com/DaveSaah/voting-system/db/api"
	"github.com/DaveSaah/voting-system/handlers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/login", handlers.Login)

	go func() {
		log.Println("Starting api server: Listening on :8000")
		_ = http.ListenAndServe(":8000", api.Init())
	}()

	log.Println("Starting main server: Listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
