package postgres

import (
	"errors"
	"fmt"
	"hash/fnv"
)

func hashFunc(s string) string {
	h := fnv.New64a()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum64())
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

	healthHash := hashFunc(voter.HealthCard)
	nameHash := hashFunc(voter.Name)

	if voterData.HealthCardHash != healthHash || voterData.NameHash != nameHash {
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
