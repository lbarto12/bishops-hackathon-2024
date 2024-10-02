package serve

import (
	"net/http"
	"votingapi/src/handlers"
)

func AddApiHandlers(mux *http.ServeMux) {
	handlers.AddVotingHandlers(mux)
	handlers.AddTabulationHandlers(mux)
}
