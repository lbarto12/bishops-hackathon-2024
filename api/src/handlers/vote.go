package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"votingapi/src/postgres"
	"votingapi/src/responses"
)

func AddVotingHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/vote", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{})
		}

		log.Printf("called with: %v", string(body))

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
