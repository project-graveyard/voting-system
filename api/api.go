package api

import (
	"log"
	"net/http"
)

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to api home\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api", &homeHandler{})
	return mux
}
