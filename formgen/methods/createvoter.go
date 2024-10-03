package methods

import (
	"formgen/postgres"
	"formgen/util"
	"log"
)

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

	_, err = tx.Exec(`
INSERT INTO voter 
    (health_card_hash, name_hash, candidate_1, candidate_2, candidate_3, has_voted) 
	VALUES ($1, $2, $3, $4, $5, $6)`,
		util.HashFunc(healthCard),
		util.HashFunc(name),
		uuids[0],
		uuids[1],
		uuids[2],
		false,
	)
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
