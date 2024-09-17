package postgres

import (
	"database/sql"
	"errors"
)

func (voter Voter) FillStats() (Voter, error) {
	db, err := Connect()
	if err != nil {
		return Voter{}, err
	}
	defer db.Close()

	voter.CanVote = false

	if err := db.QueryRow(`
	SELECT has_voted from voters where health_card = $1;
`, voter.HealthCard).Scan(&voter.HasVoted); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return voter, nil
		} else {
			return Voter{}, err
		}
	}

	voter.CanVote = true

	return voter, nil
}
