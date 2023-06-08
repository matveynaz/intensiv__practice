package main

import (
	"log"
	"net/http"
	"services/contact/internal"
	"services/contact/internal/repository"
)

func main() {
	// Create repositories
	contactRepo := repository.NewContactRepository()
	groupRepo := repository.NewGroupRepository()

	// Create the ContactService with the constructed layers
	contactService := internal.NewContactService(contactRepo, groupRepo)

	// Get the HTTP handler from the ContactService
	httpHandler := contactService.HTTPHandler

	// Set up the HTTP server
	http.HandleFunc("/contacts", httpHandler.HandleContacts)

	// Start listening for incoming requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}
