package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
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
		uid, _ := strconv.ParseInt(string([]byte(r.PathValue("uid"))[:8]), 16, 16)
		uidInt := int(uid)
		uidInt = uidInt % 100
		w.Write([]byte(strconv.Itoa(uidInt)))
	})

}
