package rest

import (
	"github.com/gorilla/mux"
	reading "github.com/justincremer/go-api/pkg/read"
)

func InitHandlers(rs reading.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/users", listUsersHandler(rs)).Methods("GET")

	return router
}
