package postgres

import (
	"errors"
)

func Vote(voter Voter, candidate string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	voterData, err := voter.FillStats()
	if err != nil {
		return err
	}

	if !voterData.CanVote {
		return errors.New("voter is not eligible to vote")
	}

	if voterData.HasVoted {
		return errors.New("voter has already voted")
	}

	_, err = db.Query(`
	UPDATE polls
	SET votes = votes + 1
	WHERE candidate = $1
`, candidate)
	if err != nil {
		return err
	}

	_, err = db.Query(`
	UPDATE voters
	SET has_voted = true
	WHERE health_card = $1
`, voter.HealthCard)
	if err != nil {
		return err
	}

	return nil
}
