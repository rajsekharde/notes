package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// GET /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Backend server running")

	log.Printf("%s %s %d OK", r.Method, r.URL.Path, http.StatusOK)
}

// POST /users
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON body into a User struct
	var newUser User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := insertUser(newUser.Name, newUser.Age)
	if err != nil {
		log.Println(err)
		return
	}

	// Hard coded insert
	// err := insertUser("RS", 21)
	// if err != nil {
	// 	log.Println(err)
	// }

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User inserted")

	log.Printf("%s %s %d CREATED", r.Method, r.URL.Path, http.StatusCreated)
}

// GET /users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := getUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Returning JSON response
	w.Header().Set("Content-Type", "application/json") // Header should be set before writing
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Returning string response
	// w.WriteHeader(http.StatusOK)
	// for _, u := range users {
	// 	fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
	// }

	log.Printf("%s %s %d OK", r.Method, r.URL.Path, http.StatusOK)
}

// GET /users/<ID>
func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id") // Get path parameters (id)

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	u, err := getUserById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)

	log.Printf("%s %s %d OK", r.Method, r.URL.Path, http.StatusOK)
}

type UserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PUT /users/{id}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse the incoming JSON body into a UserRequest struct
	var updatedUser User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = updateUser(id, updatedUser.Name, updatedUser.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// hard coded update
	// err := updateUser(2, "KA", 21)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "User updated successfully"}`)

	log.Printf("%s %s %d OK", r.Method, r.URL.Path, http.StatusOK)
}

// DELETE /users/{id}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = deleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User deleted")

	log.Printf("%s %s %d OK", r.Method, r.URL.Path, http.StatusOK)
}
