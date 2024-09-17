package postgres

func Vote(voter Voter) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query(`
	UPDATE voters
	SET has_voted = true
	WHERE health_card = $1
`, voter.HealthCard)

	return err
}
