package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // Write status code to response header

	fmt.Fprintln(w, "Hello World from Go backend!") // Write message to response body

	log.Printf("%s %s %d", r.Method, r.URL.Path, http.StatusOK) // Logging request
	/*
	r.Method- returns http method used for the request
	r.URL.Path - returns request endpoint
	http.StatusOK - returns 200
	*/
}

func main() {

	// Simple HTTP server using the default request router
	// http.HandleFunc("/", handler)
	// fmt.Println("Server running on :8080")
	// http.ListenAndServe(":8080", nil) // nil -> uses default router

	// Using custom router
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler) // Valid for Go-1.22+
	// mux.HandleFunc("/", handler) // Valid for all versions
	log.Println("Server running at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

/*
http.HandleFunc() - Maps route -> Handler
	http.HandleFunc("/", handler) - When "/" is requested, call handler()

func handler(w http.ResponseWriter, r *http.Request)
	- Handler function signature
	w - response writer (Send response)
	r - incoming request data

mux := http.NewServeMux() - Initialize a custom built in request router
mux.HandleFunc("GET /", handler):
	Define an endpoint as- "<Method> <route>" or just "<route>" for Go
	versions older than 1.22

log.Println("Server running at localhost:8000") :
	Prints message with a timestamp

http.ListenAndServer(":8080", <router>) - Starting server
	- :8080 - port number
	- <router> - Request router:
		nil - default router- http.DefaultServerMux(). Sets endpoints globally
		mux - custom router
*/