// api package handles interactions with the database
// It checks the endpoints and determines which actions to take
// depending on the HTTP request.
package api

import (
	"fmt"
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
		switch r.Method {
		case http.MethodGet:
			u.listUsers(w, r)
			return
		default:
			err := fmt.Sprintf("Error %d: Method not allowed\n", http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte(err))
			return
		}
	}

	// get access to the query string -> everything after "?"
	qs := strings.TrimPrefix(r.RequestURI, r.URL.Path+"?")

	// store query parameters as a slice -> key=val
	qp := strings.Split(qs, "&")

	params := make(map[string]string)

	// set params as (key, value) pairs
	for _, val := range qp {
		v := strings.Split(val, "=")
		params[v[0]] = v[1]
	}

	switch r.Method {
	case http.MethodGet:
		u.getUser(w, r, params["id"])
		return
	case http.MethodPost:
		u.createUser(w, r, params)
		return
	}
}

// createUser creates a new user in the database
// It handles POST requests to /api/users
func (u *usersHandler) createUser(w http.ResponseWriter, r *http.Request, data map[string]string) {
	msg := fmt.Sprintf(
		"Created a new user with:\nID: %s, Name: %s, Class:%s\n",
		data["id"], data["fname"]+" "+data["lname"], data["class"],
	)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

// listUsers retrieves all users from the database
// It handles GET requests to /api/users
func (u *usersHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("User list\n"))
	if err != nil {
		log.Fatal(err)
	}
}

// getUser retrieves a user from the database
// It handles GET requests to /api/users/{id}
func (u *usersHandler) getUser(w http.ResponseWriter, r *http.Request, id string) {
	msg := fmt.Sprintf("User %s retrieved\n", id)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/users", &usersHandler{})
	mux.Handle("/api/users/", &usersHandler{})

	return mux
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
