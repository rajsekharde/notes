package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserFetch struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	usersR, err := getUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []UserFetch
	for _, u := range usersR {
		res = append(res, UserFetch{u.id, u.name, u.age})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("%s %s %d\n", r.Method, r.URL.Path, http.StatusOK)
}