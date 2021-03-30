package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justincremer/go-api/pkg/http/rest"
)

func main() {
	port := ":8080"

	fmt.Printf("Server running on http://localhost%v", port)
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(port, router))
}
