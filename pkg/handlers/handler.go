package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justincremer/go-api/pkg/services"
)

func rootHandler() func(writer http.ResponseWriter, response *http.Request) {
	return func(writer http.ResponseWriter, response *http.Request) {
		json.NewEncoder(writer).Encode("Welcome to the money store")
	}
}

func InitHandlers(userService services.UserService) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler()).Methods("GET")
	router.HandleFunc("/users", listUsersHandler(userService)).Methods("GET")
	router.HandleFunc("/users", createUserHandler(userService)).Methods("POST")
	// router.HandleFunc("/users", updateUserHandler(userService)).Methods("PATCH")
	// router.HandleFunc("/users", deleteUserHandler(userService)).Methods("DELETE")

	return router
}
