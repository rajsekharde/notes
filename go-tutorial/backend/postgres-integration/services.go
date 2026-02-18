package main

import (
	"fmt"
	"log"
)

// Create a table in the DB
func createUsersTable() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);
	`

	_, err := db.Exec(createTableQuery) // Exec() does not return rows
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created table 'users'")
}


// Insert a row into users table
func insertUser(name string, age int) error {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`

	_, err := db.Exec(query, name, age)

	return err
}


type User struct {
	ID   int
	Name string
	Age  int
}

// Get data of all rows
func getUsers() ([]User, error) {
	rows, err := db.Query(`SELECT id, name, age FROM users`) // Query() returns rows along with error
	if err !=nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}


// Get data of user by id
func getUserById(ID int) (User, error) {
	query := `
	SELECT id, name, age FROM users WHERE id=$1
	`

	var user User
	err := db.QueryRow(query, ID).Scan(&user.ID, &user.Name, &user.Age)

	return user, err
}


// Update data of a row
func updateUser(id int, name string, age int) error {
	query := `
	UPDATE users
	SET name=$1, age=$2
	WHERE id=$3
	`

	result, err := db.Exec(query, name, age, id)
	if err != nil {
		return err
	}

	// Get number of rows changed
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Check if id exists in the db
	if rowsAffected == 0 {
		return fmt.Errorf("No user found with id %d", id)
	}

	return nil
}


// Delete a row
func deleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No user found with id %d", id)
	}

	return nil
}