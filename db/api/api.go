// api package handles interactions with the database
// It checks the endpoints and determines which actions to take
// depending on the HTTP request.
package api

import (
	"log"
	"net/http"
	"regexp"
)

// define route handlers for api
type (
	homeHandler  struct{}
	usersHandler struct{}
)

var (
	// matches default users route: /users
	userReq = regexp.MustCompile(`^/users/*$`)

	// matches users route containing id: /users/<id>
	userReqWithID = regexp.MustCompile(`^/users/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

// define a store for users
type Store interface {
	Add() error
	Get() error
	Update() error
	List() error
	Remove() error
}

// ServeHTTP for homeHandler handles the /api route
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// display a message
	_, err := w.Write([]byte("Welcome to api home\n"))
	if err != nil {
		log.Fatal(err)
	}
}

// ServeHTTP for usersHandler handles the /api/users route
func (u *usersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// display a message
	_, err := w.Write([]byte("Welcome to users route\n"))
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case r.Method == http.MethodPost && userReq.MatchString(r.URL.Path):
		u.createUser(w, r)
		return
	case r.Method == http.MethodGet && userReq.MatchString(r.URL.Path):
		u.listUsers(w, r)
		return
	case r.Method == http.MethodGet && userReqWithID.MatchString(r.URL.Path):
		u.getUser(w, r)
		return
	case r.Method == http.MethodPut && userReqWithID.MatchString(r.URL.Path):
		u.updateUser(w, r)
		return
	case r.Method == http.MethodDelete && userReqWithID.MatchString(r.URL.Path):
		u.deleteUser(w, r)
		return
	default:
		return
	}
}

func (u *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {}
func (u *usersHandler) listUsers(w http.ResponseWriter, r *http.Request)  {}
func (u *usersHandler) getUser(w http.ResponseWriter, r *http.Request)    {}
func (u *usersHandler) updateUser(w http.ResponseWriter, r *http.Request) {}
func (u *usersHandler) deleteUser(w http.ResponseWriter, r *http.Request) {}

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api", &homeHandler{})
	mux.Handle("/api/users", &usersHandler{})
	mux.Handle("/api/users/", &usersHandler{})

	return mux
}
