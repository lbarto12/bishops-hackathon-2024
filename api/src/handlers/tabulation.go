package handlers

import (
	"net/http"
	"votingapi/src/postgres"
	"votingapi/src/responses"
)

// AddTabulationHandlers Adds tabulation handlers to the server mutex
func AddTabulationHandlers(mux *http.ServeMux) {

	// Returns the current state of the polls. simply, candidates and respective number of votes
	mux.HandleFunc("GET /api/tabulation/polls", func(w http.ResponseWriter, r *http.Request) {
		data, err := postgres.GetPolls()

		if err != nil {
			responses.DoErrorResponse(w, responses.ApiResponse[any]{
				Status:  http.StatusInternalServerError,
				Message: "Error retrieving polls",
				Data:    nil,
			})
			return
		}

		responses.DoSuccessResponse(w, responses.ApiResponse[postgres.PollData]{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    data,
		})

	})

}
