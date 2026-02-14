package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Backend World!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

/*
http.HandleFunc() - Maps route -> Handler
	http.HandleFunc("/", handler) - When "/" is requested, call handler()

func handler(w http.ResponseWriter, r *http.Request)
	- Handler function signature
	w - response writer (Send response)
	r - incoming request data

http.ListenAndServer(":8080", nil) - Starting server
	- :8080 - port number
	- nil - default router
*/