package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func DoSuccessResponse[T any](w http.ResponseWriter, data ApiResponse[T]) {
	w.WriteHeader(data.Status)
	w.Header().Add("Content-Type", "application/json")
	message, err := json.Marshal(data)
	if err != nil {
		DoErrorResponse(
			w,
			ApiResponse[T]{
				Status:  http.StatusInternalServerError,
				Message: "Failed to marshal response data",
			},
		)
	}
	_, err = w.Write(message)
	if err != nil {
		log.Printf("Error writing api success response: %s", err)
		return
	}
}

func DoErrorResponse[T any](w http.ResponseWriter, response ApiResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	if response.Message != "" {
		response.Message = http.StatusText(response.Status) + ": " + response.Message
	} else {
		response.Message = http.StatusText(response.Status)
	}
	returnMessage, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshalling error response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(`{"message": "error marshalling error response"}`)); err != nil {
			log.Printf("Error writing error response: %v", err)
		}
		return
	}
	w.WriteHeader(response.Status)
	_, err = w.Write(returnMessage)
	if err != nil {
		log.Printf("Error writing api error response: %s", err)
		return
	}
}
