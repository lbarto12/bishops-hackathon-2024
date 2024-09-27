package postgres

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

/*
	hashFunc hashes the input string using bcrypt.

no salting needed when using bcrypt  since it automatically generates a salt and returns a hashed string.
*/
func hashFunc(input string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Compare the plain input with the stored hash.
func compareHash(hash, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
}

var HasVotedError = errors.New("user has already voted")
var InvalidVoteError = errors.New("invalid vote")

func Vote(voter VoterRequest, candidate string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	var voterData VoterData
	err = db.Get(&voterData, "SELECT * FROM voter WHERE candidate_1 = $1 OR candidate_2 = $1 OR candidate_3 = $1", candidate)
	if err != nil {
		return err
	}

	if voterData.HasVoted {
		return HasVotedError
	}

	// Comparing the stored hash with the input voter data.
	if compareHash(voterData.HealthCardHash, voter.HealthCard) != nil || compareHash(voterData.NameHash, voter.Name) != nil {
		return InvalidVoteError
	}

	var candidateId int
	if voterData.Candidate1 == candidate {
		candidateId = 1
	} else if voterData.Candidate2 == candidate {
		candidateId = 2
	} else if voterData.Candidate3 == candidate {
		candidateId = 3
	}

	tx, err := db.Beginx()

	if err != nil {
		return err
	}

	query, err := tx.NamedQuery(`UPDATE voter SET has_voted = true WHERE id = :id;`, voterData)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = query.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE polls SET votes = votes + 1 WHERE candidate = $1;`, candidateId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
