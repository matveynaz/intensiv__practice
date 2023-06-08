package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Connection parameters
	host := "localhost"
	port := 5432
	user := "your_user"
	password := "your_password"
	dbname := "your_dbname"

	// Establish a connection
	conn, err := connectToPostgreSQL(host, port, user, password, dbname)
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

func connectToPostgreSQL(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	// Connection string format
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
