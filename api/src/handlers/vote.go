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

func AddVotingHandlers(mux *http.ServeMux) {
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
				Message: "Vote not Registered: " + err.Error(),
			})
			return
		}

		responses.DoSuccessResponse(w, responses.ApiResponse[any]{
			Status:  http.StatusCreated,
			Message: "Vote Created",
		})

	})

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
		fmt.Println(fmt.Sprintf("%x", ha[:2]))

		io.WriteString(w, fmt.Sprintf("%x", ha[:2]))
	})

}
