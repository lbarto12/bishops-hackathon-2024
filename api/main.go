package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"votingapi/src/postgres"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment: %v", err)
		return
	}

	postgres.Vote(postgres.Voter{
		Name:       "Liam Barrack",
		HealthCard: "123-456-789",
		Vote:       "Liam",
	})

	frontendPath, ok := os.LookupEnv("FRONTEND_URL")
	if !ok {
		log.Fatal("Environment variable `FRONTEND_URL` not set:")
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{frontendPath},
		AllowCredentials: true,
		// Debug: true,
	})

	mux := http.NewServeMux()

	handler := c.Handler(mux)

	log.Printf("API: Listening on %s...\n", apiUrl)
	err = http.ListenAndServe(apiUrl, handler)
	if err != nil {
		log.Fatal(err)
	}

}
