package handlers

import (
	"net/http"
	"votingapi/src/postgres"
	"votingapi/src/responses"
)

func AddTabulationHandlers(mux *http.ServeMux) {

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
