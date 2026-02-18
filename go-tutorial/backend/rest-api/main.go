package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Create all REST API endpoints- GET, POST, PUT, PATCH, DELETE using net/http package

*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Test object (temp database)
var user1 = User{"RSD", 21}

// Single handler for all 'user' api endpoints
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getUsers(w, r)

	case http.MethodPost:
		createUser(w, r)

	case http.MethodPut:
		updateUser(w, r)

	case http.MethodPatch:
		patchUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// GET method
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]User{
		{"RS", 21},
		{"SD", 25},
	})
}

// POST method
func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	fmt.Println("Created:", u)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// PUT method
func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User fully updated (PUT)")
}

// PATCH method
func patchUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User partiallu updated (PATCH)")
}

// DELETE method
func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User deleted")
}

// GET handler
func getUser(w http.ResponseWriter, r *http.Request) {
	// user := User{"RS", 21}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user1)
}

func main() {
	// http.HandleFunc("/user", getUser)

	// Mapping all "/users" endpoints to the common handler
	http.HandleFunc("/users", usersHandler)

	// Printing a message before starting server
	fmt.Println("Server running on port 8080")

	// Starting server on port 8080
	http.ListenAndServe(":8080", nil)
}

/*
Additions:

Use mux := http.NewServeMux() as a custom router instead of default global router
	+ mux.HandleFunc("GET /users", usersHandler) to avoid method checks later on

Use log.Println for logging with timestamps
*/
