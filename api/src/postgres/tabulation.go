package postgres

// TODO: @Jerome, Dangerous, need to discuss.
func GetPolls() (PollData, error) {

	db, err := Database()
	if err != nil {
		return PollData{}, err
	}

	//defer db.Close() // temp solution to make the connection open

	var data PollData

	rows, err := db.Query("SELECT * FROM polls;") // candidate is an integer it needs to be changed
	if err != nil {
		return PollData{}, err
	}

	//defer rows.Close()  // temp solution to make the connection open

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
