package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", getUsersHandler)

	log.Println("Server running at localhost:8000")
	http.ListenAndServe(":8000", mux)
}