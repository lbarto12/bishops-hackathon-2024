package main

import (
	"formgen/methods"
	"formgen/postgres"
	"formgen/util"
	"github.com/joho/godotenv"
	"log"
)

// This populates the database with random info on people, and generates QR codes per candidate per voter
// that are stored in the `out` folder
func main() {
	util.Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	err = postgres.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = methods.ResetDB()

	err = methods.GenerateStaticVoters()
	if err != nil {

		log.Fatal(err)
		return
	}
}
