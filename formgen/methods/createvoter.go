package methods

import (
	"crypto/sha256"
	"fmt"
	"formgen/postgres"
	"log"
)

func createRegVoter(name string, healthCard string, uuids []string, verification []string) error {
	db, err := postgres.Database()
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = tx.Exec(`
INSERT INTO voter_reg 
    (health_card, name, candidate_1, candidate_2, candidate_3, can_verify_1, can_verify_2, can_verify_3) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
		healthCard,
		name,
		uuids[0],
		uuids[1],
		uuids[2],
		verification[0],
		verification[1],
		verification[2],
	)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	return tx.Commit()
}

// CreateVoter creates a single voter in the Postgres database with the uuids and hashes received
func CreateVoter(name string, healthCard string, uuids []string) error {

	db, err := postgres.Database()
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
		return err
	}

	can1 := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%x", sha256.Sum256([]byte(name+healthCard+uuids[0]))))))
	can2 := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%x", sha256.Sum256([]byte(name+healthCard+uuids[1]))))))
	can3 := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%x", sha256.Sum256([]byte(name+healthCard+uuids[2]))))))

	_, err = tx.Exec(`
	INSERT INTO voter
	(candidate_1, candidate_2, candidate_3, has_voted)
	VALUES ($1, $2, $3, $4)`, can1, can2, can3, false)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	return nil
}
