package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"votingapi/src/postgres"
	"votingapi/src/responses"
)

// AddVotingHandlers adds voting handlers to server mutex
func AddVotingHandlers(mux *http.ServeMux) {

	// registers a single vote. Accepts VoteRequest as the request type and requests
	// voting functionality from postgres
	mux.HandleFunc("POST /api/vote", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "Request Unreadable: " + err.Error(),
			})
			return
		}

		var request VoteRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status:  http.StatusBadRequest,
				Message: "Request Unreadable: " + err.Error(),
			})
			return
		}

		err = postgres.Vote(request.Data)
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
				Message: "Vote not Registered: " + err.Error(),
			})
			return
		}

		responses.DoSuccessResponse(w, responses.ApiResponse[any]{
			Status:  http.StatusCreated,
			Message: "Vote Created",
		})

	})

	// Generates and returns the candidate validation code so that client users can compare
	// it to the candidate they are trying to vote for. For anonymity purposes
	mux.HandleFunc("GET /api/getint/{uid}", func(w http.ResponseWriter, r *http.Request) {

		salt, ok := os.LookupEnv("CANDIDATE_SALT")
		if !ok {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status: http.StatusInternalServerError,
			})
			return
		}

		id := r.PathValue("uid") + salt

		ha := sha256.Sum256([]byte(id))

		io.WriteString(w, fmt.Sprintf("%x", ha[:2]))
	})

}
