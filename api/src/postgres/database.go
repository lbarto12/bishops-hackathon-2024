package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var database *sqlx.DB

func Init() error {
	var err error
	database, err = Connect()
	if err == nil {
		log.Println("Successfully connected to database")
	}
	return err
}

func Database() (*sqlx.DB, error) {
	return database, database.Ping()
}
