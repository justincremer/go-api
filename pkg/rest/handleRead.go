package rest

import (
	"encoding/json"
	"net/http"

	reading "github.com/justincremer/go-api/pkg/read"
)

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to the money store")
	}
}

func listUsersHandler(rs reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := rs.ListUsers()
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(users)
	}
}
