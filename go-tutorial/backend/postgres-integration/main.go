package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initDB()
	createUsersTable()
	defer db.Close() // Close database connection when application exits

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /users/{id}", getUserByIdHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	log.Printf("Server running at localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
