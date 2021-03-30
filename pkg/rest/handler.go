package rest

import (
	"github.com/gorilla/mux"
	reading "github.com/justincremer/go-api/pkg/read"
	writing "github.com/justincremer/go-api/pkg/write"
)

func InitHandlers(readingService reading.Service, writingService writing.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/users", listUsersHandler(readingService)).Methods("GET")
	router.HandleFunc("/users", createUserHandler(writingService)).Methods("POST")

	return router
}
