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
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "Request Unreadable",
			})
			return
		}

		log.Printf("called with: %v", string(body))

		var request VoteRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "Request Unreadable",
			})
			return
		}

		err = postgres.Vote(request.Voter, request.Candidate)
		if err != nil {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status:  http.StatusInternalServerError,
				Message: "Vote not Registered",
			})
			return
		}

		responses.DoSuccessResponse(w, responses.ApiResponse[any]{
			Status:  http.StatusCreated,
			Message: "Vote Created",
		})

	})
}
