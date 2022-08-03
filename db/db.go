package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() string {
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr.Error()
	}
	return "Connected!"
}
