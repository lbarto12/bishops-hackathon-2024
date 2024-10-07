package postgres

// TODO: @Jerome, Dangerous, need to discuss.
func GetPolls() (PollData, error) {

	db, err := Database()
	if err != nil {
		return PollData{}, err
	}

	var data PollData

	rows, err := db.Query("SELECT * FROM polls ORDER BY candidate;")
	if err != nil {
		return PollData{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var candidateData CandidatePollData
		err := rows.Scan(&candidateData.Candidate, &candidateData.Votes)
		if err != nil {
			return PollData{}, err
		}
		data.Polls = append(data.Polls, candidateData)
	}

	return data, nil
}
