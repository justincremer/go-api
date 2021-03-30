package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	reading "github.com/justincremer/go-api/pkg/read"
	"github.com/justincremer/go-api/pkg/rest"
	"github.com/justincremer/go-api/pkg/storage"
)

func main() {
	const port string = ":3000"

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatal("Error while setting up storage:", err)
	}

	rs := reading.NewService(r)

	fmt.Printf("Server running on http://localhost%v", port)
	var router *mux.Router = rest.InitHandlers(rs)

	log.Fatal(http.ListenAndServe(port, router))
}
