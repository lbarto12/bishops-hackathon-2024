package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// database Package private, but globally accessible reference to postgres
var database *sqlx.DB

// Init calls the connection and sets relevant properties of the database, returns failure
func Init() error {
	var err error
	database, err = Connect()
	if err == nil {
		log.Println("Successfully connected to database")
	}
	database.SetMaxOpenConns(75)
	return err
}

// Database simply returns the current instance of the connection pool to caller, along with an error if
// The database cannot be reached
func Database() (*sqlx.DB, error) {
	return database, database.Ping()
}
