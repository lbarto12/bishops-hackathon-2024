package postgres

import (
	"database/sql"
	"fmt"
	"os"
)

func GetCredentials() Credentials {
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		panic("POSTGRES_HOST environment variable not set")
	}

	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		panic("POSTGRES_PORT environment variable not set")
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		panic("POSTGRES_USER environment variable not set")
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		panic("POSTGRES_PASSWORD environment variable not set")
	}

	dbname, ok := os.LookupEnv("POSTGRES_DBNAME")
	if !ok {
		panic("POSTGRES_DBNAME environment variable not set")
	}

	return Credentials{
		Host:     host,
		Port:     port,
		User:     user,
		Pass:     password,
		Database: dbname,
	}
}

func Connect() (*sql.DB, error) {

	credentials := GetCredentials()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		credentials.Host,
		credentials.Port,
		credentials.User,
		credentials.Pass,
		credentials.Database,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
