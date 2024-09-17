package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"votingapi/src/postgres"
)

func AddVotingHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/vote", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		var request VoteRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		err = postgres.Vote(request.Voter, request.Candidate)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		io.WriteString(w, "Vote Successful!")

	})
}
