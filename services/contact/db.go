package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgreSQL(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
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
