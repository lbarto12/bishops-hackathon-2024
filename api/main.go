package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"
	"votingapi/src/handlers"
	"votingapi/src/postgres"
)

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

	handlers.AddVotingHandlers(mux)

	handler := c.Handler(mux)

	// Demo call
	go func() {
		time.Sleep(time.Second * 3)

		// TODO: implement encryption on DB, currently just plain text :'(
		req, err := json.Marshal(handlers.VoteRequest{
			Voter: postgres.Voter{
				HealthCard: "123-456-789",
			},
			Candidate: "Liam",
		})

		resp, err := http.Post(fmt.Sprintf("http://%s/api/vote", apiUrl), "application/json", bytes.NewBuffer(req))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Printf("Response status: %v\n", resp.Body) //TODO: implement proper response codes for front-end

	}()

	log.Printf("API: Listening on %s...\n", apiUrl)
	err = http.ListenAndServe(apiUrl, handler)
	if err != nil {
		log.Fatal(err)
	}

}
