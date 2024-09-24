package handlers

import (
	"encoding/json"
	"errors"
	"io"
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
			if errors.Is(err, postgres.HasVotedError) {
				responses.DoSuccessResponse(w, responses.ApiResponse[any]{
					Status:  http.StatusOK,
					Message: "Vote Created",
				})
				return
			}

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
