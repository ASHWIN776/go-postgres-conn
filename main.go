package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect to DB
	conn, err := sql.Open("pgx", fmt.Sprintf("host=localhost port=5432 dbname=go-postgres user=%s password=%s", username, pass))

	if err != nil {
		log.Fatal("failed to connect: ", err)
	}

	// Test Connection
	err = conn.Ping()

	if err != nil {
		log.Fatal("failed to ping database")
	}

	log.Println("Connected")

	// Defer - close connection
	defer conn.Close()

	// Get all rows from tasql

	// Insert a row

	// Get all rows again

	// Update a row

	// Get all rows again

	// Get row by id

	// Delete row by id

	// Get all rows again
}
