package postgres

type Credentials struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type VoterData struct {
	Id         int    `db:"id"`
	HasVoted   bool   `db:"has_voted"`
	Candidate1 string `db:"candidate_1"`
	Candidate2 string `db:"candidate_2"`
	Candidate3 string `db:"candidate_3"`
}

type VoterRequest struct {
	Data string `json:"data"`
}

type CandidatePollData struct {
	Candidate int `json:"candidate"`
	Votes     int `json:"votes"`
}

type PollData struct {
	Polls []CandidatePollData `json:"polls"`
}
