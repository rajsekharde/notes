package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)


func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello")
	log.Printf("%s %s %d", r.Method, r.URL.Path, http.StatusOK)
}


type UserJson struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Password hashing
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

type signUpResponse struct {
	Token string `json:"access_token"`
}

// Handle Sign Up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var u UserJson
	// Extract the JSON body from request and store in u
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCount += 1
	// Get hashed password
	passwordHash, err := hashPassword(u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = userCount

	// Create new user struct
	newUser := User{
		u.ID,
		u.Email,
		passwordHash,
	}

	// Add new user to database
	DB = append(DB, newUser)

	fmt.Printf("ID: %d\nEmail: %s\nPassword Hash: %s\n", u.ID, u.Email, passwordHash)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User Registered")

	log.Printf("%s %s %d", r.Method, r.URL.Path, http.StatusOK)
}

// Verify password by hashing and comparing to stored hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// JWT secret key
var jwtKey = []byte("secret_key")

// Generate JWT
func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}


// Handle Log In
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u UserJson
	err2 := json.NewDecoder(r.Body).Decode(&u)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	// Check if credentials are correct
	verified := false
	for _, ud := range DB {
		if ud.Email == u.Email && CheckPasswordHash(u.Password, ud.Password) {
			verified = true
			break
		}
	}

	if verified == false {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Incorrect Credentials")
		return
	}

	// Generate a JWT
	newToken, _ := generateJWT(u.Email)
	resp := signUpResponse{
		Token: newToken,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode JWT into response
	json.NewEncoder(w).Encode(resp)

	log.Printf("%s %s %d", r.Method, r.URL.Path, http.StatusOK)
}

type GetUsersResp struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// Protected route
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from request header
	tokenString := r.Header.Get("Authorization")
	// Strip "Bearer" prefix
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Not Authorised")
		return
	}

	// Get all users and store in a struct slice
	var resp []GetUsersResp
	for _, u := range DB {
		user := GetUsersResp {
			ID: u.ID,
			Email: u.Email,
			PasswordHash: u.Password,
		}
		resp = append(resp, user)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode slice struct to response body
	json.NewEncoder(w).Encode(&resp)

	log.Printf("%s %s %d", r.Method, r.URL.Path, http.StatusOK)
}
