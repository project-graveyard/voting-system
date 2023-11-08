// api package handles interactions with the database
// It checks the endpoints and determines which actions to take
// depending on the HTTP request.
package api

import (
	"log"
	"net/http"
	"strings"
)

// define route handlers for api
type (
	usersHandler struct{}
)

// ServeHTTP for usersHandler handles the /api/users route
func (u *usersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check if request includes query parameters
	if ok := strings.Contains(r.URL.RequestURI(), "?"); !ok {
		log.Printf("%s: no query data\n", r.Method)

		switch r.Method {
		case http.MethodGet:
			u.listUsers(w, r)
			return
		case http.MethodPost:
			u.createUser(w, r)
			return
		}
	}

	// if there are query parameters,
	log.Printf("%s: query data included\n", r.Method)
	params := make(map[string]any)

	// get access to the query string -> everything after ?
	query_str := strings.TrimPrefix(r.RequestURI, r.URL.Path+"?")
	log.Printf("%s\n", query_str)

	// store query parameters as a slice -> key=val
	query_params := strings.Split(query_str, "&")
	log.Printf("%v\n", query_params)

	for _, val := range query_params {
		v := strings.Split(val, "=")
		params[v[0]] = v[1]
	}

	log.Printf("%v\n", params)

	switch r.Method {
	case http.MethodGet:
		u.getUser(w, r)
		return
	}
}

// case r.Method == http.MethodPost:
//
//	u.createUser(w, r)
//	return
//
//
// case r.Method == http.MethodPut && userReqWithID.MatchString(r.URL.Path):
// 	u.updateUser(w, r)
// 	return
// case r.Method == http.MethodDelete && userReqWithID.MatchString(r.URL.Path):
// 	u.deleteUser(w, r)
// 	return
// default:
// 	return
// }
// }

func (u *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("User created\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func (u *usersHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("User list\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func (u *usersHandler) getUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("User data retrieved\n"))
	if err != nil {
		log.Fatal(err)
	}
}

// func (u *usersHandler) updateUser(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte("User data updated\n"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func (u *usersHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte("User data deleted\n"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/users", &usersHandler{})
	mux.Handle("/api/users/", &usersHandler{})

	return mux
}
