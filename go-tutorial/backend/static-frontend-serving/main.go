package main

import (
	"fmt"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This API was served using Go")
}

func main() {
	// Create handler & define folder to serve
	fs := http.FileServer(http.Dir("./static"))

	// Map route to handler
	http.Handle("/", fs)

	// Create api endpoint (additional)
	http.HandleFunc("/api/", apiHandler)

	// Run the server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

/*
http.FileServer() - Creates a handler that:
	- Finds requested file
	- Serves correct MIME type
	- Handles caching basics

http.Dir("./static") - Defines folder to serve

http.Handle("/", fs) - Maps "/" route to 'fs' handler

GET /style.css -> ./static/style.css
*/
