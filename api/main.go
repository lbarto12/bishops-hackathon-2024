package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"votingapi/src/postgres"
	"votingapi/src/serve"
)

// This is the API. Manages requests between the client and postgres
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment: %v\n", err)
		return
	}

	frontendPath, ok := os.LookupEnv("FRONTEND_URL")
	if !ok {
		log.Fatal("Environment variable `FRONTEND_URL` not set:")
	}

	tabulationPath, ok := os.LookupEnv("TABULATION_URL")
	if !ok {
		log.Fatal("Environment variable `TABULATION_URL` not set:")
	}

	apiHost, ok := os.LookupEnv("API_HOST")
	if !ok {
		log.Fatal("Environment variable `API_HOST` not set:")
	}

	apiPort, ok := os.LookupEnv("API_PORT")
	if !ok {
		log.Fatal("Environment variable `API_PORT` not set:")
	}

	apiUrl := fmt.Sprintf("%s:%s", apiHost, apiPort)

	// Adds cors headers for exclusively the frontend and tabulation websites
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{frontendPath, tabulationPath},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		// Debug: true,
	})

	// Init Postgres
	err = postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Create and serve server mutex
	mux := http.NewServeMux()

	serve.AddApiHandlers(mux)

	handler := c.Handler(mux)

	log.Printf("API: Listening on %s...\n", apiUrl)
	log.Fatal(http.ListenAndServe(apiUrl, handler))
}
