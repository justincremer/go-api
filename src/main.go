package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/justincremer/go-api/pkg/handlers"
	"github.com/justincremer/go-api/pkg/services"

	"github.com/justincremer/go-api/pkg/storage"
)

func main() {
	const port string = ":8080"

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatal("Error while setting up storage:", err)
	}

	userService := services.NewUserService(r)

	fmt.Printf("Server running on http://localhost%v", port)
	var router *mux.Router = handlers.InitHandlers(userService)

	log.Fatal(http.ListenAndServe(port, router))
}
