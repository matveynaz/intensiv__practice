package main

import (
	"fmt"
	"services/contact/database"
)

func main() {
	// Connection parameters
	host := "localhost"
	port := 5432
	user := "your_user"
	password := "your_password"
	dbname := "your_dbname"

	// Establish a connection
	conn, err := database.ConnectToPostgreSQL(host, port, user, password, dbname)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	defer conn.Close()

	// Use the connection for database operations
	// ...

	// Example: Query the database
	rows, err := conn.Query("SELECT * FROM your_table")
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return
	}
	defer rows.Close()

	// Process the query results
	// ...
}
