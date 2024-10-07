package postgres

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// HasVotedError Error definitions for returns and comparisons
var HasVotedError = errors.New("user has already voted")
var InvalidVoteError = errors.New("invalid vote")

// Vote Registers the vote, accepts a voter token - which is the hashed result from the client.
// Double hashes it and compares it to the public candidate vote table 'voter'.
// Checks if voter has already voted and fails if so.
func Vote(token string) error {
	tokenHashByte := sha256.Sum256([]byte(token))
	tokenHash := fmt.Sprintf("%x", tokenHashByte[:])

	db, err := Database()
	if err != nil {
		return err
	}

	var voterData VoterData
	err = db.Get(&voterData, "SELECT * FROM voter WHERE candidate_1 = $1 OR candidate_2 = $1 OR candidate_3 = $1", tokenHash)
	if err != nil {
		return err
	}

	if voterData.HasVoted {
		return HasVotedError
	}

	// This would be cleaner with real candidate names, ints used for ordering and simplicity in demo
	var candidateId int
	if voterData.Candidate1 == tokenHash {
		candidateId = 1
	} else if voterData.Candidate2 == tokenHash {
		candidateId = 2
	} else if voterData.Candidate3 == tokenHash {
		candidateId = 3
	}

	tx, err := db.Beginx()

	if err != nil {
		return err
	}

	// Set voter has_voted
	query, err := tx.NamedQuery(`UPDATE voter SET has_voted = true WHERE id = :id;`, voterData)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = query.Close()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Update polls
	_, err = tx.Exec(`UPDATE polls SET votes = votes + 1 WHERE candidate = $1;`, candidateId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}
