package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	// Create a request to pass to handler
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (this implements http.ResponseWriter)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsersHandler)

	// Call the handler directly
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var actualUsers []UserFetch
	if err := json.NewDecoder(rr.Body).Decode(&actualUsers); err != nil {
		t.Fatal("Could not decode response")
	}

	if len(actualUsers) != 3 {
		t.Errorf("Expected 3 users, got %d", len(actualUsers))
	}
}
