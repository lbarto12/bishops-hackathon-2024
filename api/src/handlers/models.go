package handlers

import "votingapi/src/postgres"

type VoteRequest struct {
	Voter     postgres.VoterRequest `json:"voter"`
	Candidate string                `json:"candidate"`
	Success   bool                  `json:"success,omitempty"`
}
