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

	var query string

	// Insert a row
	query = `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Jack", "Brown")

	if err != nil {
		log.Fatal("could not insert row")
	}
	log.Println("Inserted row")

	// Get all rows again
	err = getRows(conn)
	if err != nil {
		log.Fatal("failed getRows()", err)
	}

	// Update a row
	query = `update users set first_name=$1 where first_name=$2`
	res, err := conn.Exec(query, "Johnny", "John")

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := res.RowsAffected()
	log.Println("Updated rows - ", rowsAffected)

	// Get all rows again
	err = getRows(conn)
	if err != nil {
		log.Fatal("failed getRows()", err)
	}

	// Get row by id
	query = `select * from users where first_name=$1`
	row := conn.QueryRow(query, "John")
	var first_name, last_name string
	var id int

	err = row.Scan(&id, &first_name, &last_name)

	if err != nil {
		log.Println("could not get row")
	} else {
		log.Printf("Returned Row: id - %d, %s %s\n", id, first_name, last_name)
	}

	// Delete row by id
	query = `delete from users where id=$1`
	res, err = conn.Exec(query, 1)

	if err != nil {
		log.Fatal("could not delete from users ", err)
	}

	rowsAffected, _ = res.RowsAffected()
	log.Println("Deleted rows - ", rowsAffected)

	// Get all rows again
	err = getRows(conn)
	if err != nil {
		log.Fatal("failed getRows()", err)
	}
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

	// Catching the error from rows.Next()
	if err = rows.Err(); err != nil {
		log.Fatal("error of rows.Next() ", err)
	}

	log.Println("--------------------------------------------------")
	return nil
}
