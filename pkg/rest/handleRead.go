package rest

import (
	"encoding/json"
	"net/http"

	reading "github.com/justincremer/go-api/pkg/read"
)

func welcomeHandler() func(writer http.ResponseWriter, response *http.Request) {
	return func(writer http.ResponseWriter, response *http.Request) {
		json.NewEncoder(writer).Encode("Welcome to the money store")
	}
}

func listUsersHandler(readingService reading.Service) func(writer http.ResponseWriter, response *http.Request) {
	return func(writer http.ResponseWriter, response *http.Request) {
		users, err := readingService.ListUsers()
		if err != nil {
			http.Error(writer, "Error fetching users", http.StatusInternalServerError)
		}

		json.NewEncoder(writer).Encode(users)
	}
}
