package rest

import (
	"encoding/json"
	"net/http"

	writing "github.com/justincremer/go-api/pkg/write"
)

func createUserHandler(writingService writing.Service) func(writer http.ResponseWriter, response *http.Request) {
	return func(writer http.ResponseWriter, response *http.Request) {
		var newUser writing.User

		if err := json.NewDecoder(response.Body).Decode(&newUser); err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		id, err := writingService.CreateUser(newUser)
		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		newUser.Id = id
		json.NewEncoder(writer).Encode(newUser)
	}
}
