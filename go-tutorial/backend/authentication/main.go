package main

import (
	"log"
	"net/http"
)

type User struct {
	ID int
	Email string
	Password string
}

var DB = []User{}
var userCount = 0

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", RootHandler)
	mux.HandleFunc("POST /signup", SignUpHandler)
	mux.HandleFunc("POST /login", LoginHandler)
	mux.HandleFunc("GET /users", GetUsersHandler)

	log.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}