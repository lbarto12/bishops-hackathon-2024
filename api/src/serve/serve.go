package serve

import (
	"net/http"
	"votingapi/src/handlers"
)

// AddApiHandlers adds all relevant handlers to the server mux
func AddApiHandlers(mux *http.ServeMux) {
	handlers.AddVotingHandlers(mux)
	handlers.AddTabulationHandlers(mux)
}
