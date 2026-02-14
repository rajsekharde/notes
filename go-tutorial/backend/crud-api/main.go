package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

/*
Create all REST API endpoints- GET, POST, PUT, PATCH, DELETE using net/http package

*/

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

var user1 = User{"RSD", 21}

// GET endpoint
func getUser(w http.ResponseWriter, r *http.Request) {
	// user := User{"RS", 21}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user1)
}

func main() {
	http.HandleFunc("/user", getUser)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}