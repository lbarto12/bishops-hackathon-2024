package postgres

type Credentials struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type Voter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	HealthCard string `json:"health_card"`
	HasVoted   bool   `json:"has_voted"`
	Vote       string
	CanVote    bool
}
