package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect to DB
	conn, err := sql.Open("pgx", fmt.Sprintf("host=localhost port=5432 dbname=go-postgres user=%s password=%s", username, pass)) // Insert usernme and password

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

	// Get all rows from table
	err = getRows(conn)
	if err != nil {
		log.Fatal("failed getRows()", err)
	}

	// Insert a row

	// Get all rows again

	// Update a row

	// Get all rows again

	// Get row by id

	// Delete row by id

	// Get all rows again
}

func getRows(conn *sql.DB) error {
	// Executing the Select * query
	rows, err := conn.Query("select * from users")

	if err != nil {
		log.Fatal("could not return rows", err)
		return err
	}

	// Closing *sql.Rows
	defer rows.Close()
	var first_name, last_name string
	var id int

	log.Println("Getting all rows --------------------------------")
	for rowId := 1; rows.Next(); rowId++ {
		err := rows.Scan(&id, &first_name, &last_name)

		if err != nil {
			log.Fatal("could not read record")
			return err
		}

		log.Printf("Row %d: %d %s %s\n", rowId, id, first_name, last_name)
	}

	log.Println("--------------------------------------------------")
	return nil
}
