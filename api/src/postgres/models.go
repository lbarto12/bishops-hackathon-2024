package postgres

type Credentials struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type VoterData struct {
	Id             int    `db:"id""`
	NameHash       string `db:"name_hash"`
	HealthCardHash string `db:"health_card_hash"`
	HasVoted       bool   `db:"has_voted""`
	Candidate1     string `db:"candidate1"`
	Candidate2     string `db:"candidate2"`
	Candidate3     string `db:"candidate3"`
}

type VoterRequest struct {
	Name       string `json:"name"`
	HealthCard string `json:"health_card"`
}
