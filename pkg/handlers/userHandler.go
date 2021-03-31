package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/justincremer/go-api/pkg/schemas"
	"github.com/justincremer/go-api/pkg/services"
)

func listUsersHandler(userService services.UserService) func(writer http.ResponseWriter, response *http.Request) {
	return func(writer http.ResponseWriter, response *http.Request) {
		users, err := userService.ListUsers()
		if err != nil {
			http.Error(writer, "Error fetching users", http.StatusInternalServerError)
		}

		json.NewEncoder(writer).Encode(users)
	}
}

func createUserHandler(userService services.UserService) func(writer http.ResponseWriter, response *http.Request) {
	return func(writer http.ResponseWriter, response *http.Request) {
		var newUser schemas.User

		if err := json.NewDecoder(response.Body).Decode(&newUser); err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		id, err := userService.CreateUser(newUser)
		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		newUser.Id = id
		json.NewEncoder(writer).Encode(newUser)
	}
}

// func updateUserHandler(userService services.UserService) func(writer http.ResponseWriter, response *http.Request) {
// 	return func(writer http.ResponseWriter, response *http.Request) {
// 		var newUser schemas.User

// 		if err := json.NewDecoder(response.Body).Decode(&newUser); err != nil {
// 			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		id, err := userService.CreateUser(newUser)
// 		if err != nil {
// 			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		newUser.Id = id
// 		json.NewEncoder(writer).Encode(newUser)
// 	}
// // }

// func deleteUserHandler(userService services.UserService) func(writer http.ResponseWriter, response *http.Request) {
// 	return func(writer http.ResponseWriter, response *http.Request) {
// 		var newUser schemas.User

// 		if err := json.NewDecoder(response.Body).Decode(&newUser); err != nil {
// 			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		id, err := userService.CreateUser(newUser)
// 		if err != nil {
// 			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		newUser.Id = id
// 		json.NewEncoder(writer).Encode(newUser)
// 	}
// }
