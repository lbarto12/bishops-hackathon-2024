package postgres

// Credentials A struct for storing postgres credentials
type Credentials struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// VoterData Internal struct for modelling voter table queries
type VoterData struct {
	Id         int    `db:"id"`
	HasVoted   bool   `db:"has_voted"`
	Candidate1 string `db:"candidate_1"`
	Candidate2 string `db:"candidate_2"`
	Candidate3 string `db:"candidate_3"`
}

// PollData Internal struct for modelling poll table queries
type PollData struct {
	Polls []CandidatePollData `json:"polls"`
}

// CandidatePollData Internal sub-struct for modelling poll table queries
type CandidatePollData struct {
	Candidate int `json:"candidate"`
	Votes     int `json:"votes"`
}
