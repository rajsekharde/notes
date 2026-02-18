package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const PORT = 8000

var db *sql.DB

func initDB() {
	// Connection parameters
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "db"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to postgres DB")
}