package handlers

import "votingapi/src/postgres"

type VoteRequest struct {
	Voter     postgres.Voter `json:"voter"`
	Candidate string         `json:"candidate"`
}
