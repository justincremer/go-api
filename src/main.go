package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	reading "github.com/justincremer/go-api/pkg/read"
	"github.com/justincremer/go-api/pkg/rest"
	"github.com/justincremer/go-api/pkg/storage"
	writing "github.com/justincremer/go-api/pkg/write"
)

func main() {
	const port string = ":8080"

	store, err := storage.SetupStorage()
	if err != nil {
		log.Fatal("Error while setting up storage:", err)
	}

	readingService := reading.NewService(store)
	writingService := writing.NewService(store)

	fmt.Printf("Server running on http://localhost%v", port)
	var router *mux.Router = rest.InitHandlers(readingService, writingService)

	log.Fatal(http.ListenAndServe(port, router))
}
