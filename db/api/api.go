// api package handles interactions with the database
// It checks the endpoints and determines which actions to take
// depending on the HTTP request.
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DaveSaah/voting-system/db/models"
)

// define route handlers for api
type (
	usersHandler struct{}
)

// ServeHTTP for usersHandler handles the /api/users route
func (u *usersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check if request includes query parameters
	switch {
	// r.ContentLength is 0 when the r.Body is empty
	case r.Method == http.MethodGet && r.ContentLength == 0:
		u.listUsers(w, r)
		return
	case r.Method == http.MethodGet:
		u.getUser(w, r)
		return
	case r.Method == http.MethodPost:
		u.createUser(w, r)
		return
	default:
		err := fmt.Sprintf("Error %d: Method not allowed\n", http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(err))
		return
	}
}

// createUser creates a new user in the database
// It handles POST requests to /api/users
// Returns error if json data is not parsed correctly
func (u *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	// decode json data
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("New user created with ID: %s", user.ID)
}

// listUsers retrieves all users from the database
// It handles GET requests to /api/users
// Returns a json response containing the list of all users
func (u *usersHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	users = append(users, models.Users{
		Fname: "Dave",
		Lname: "Saah",
		Class: 2025,
		ID:    "72522025",
	})

	users = append(users, models.Users{
		Fname: "John",
		Lname: "Doe",
		Class: 2025,
		ID:    "72522026",
	})

	// return a json response
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(res)
}

// getUser retrieves a user from the database
// It handles GET requests to /api/users/{id}
func (u *usersHandler) getUser(w http.ResponseWriter, r *http.Request) {
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
