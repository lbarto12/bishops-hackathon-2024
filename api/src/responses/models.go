package responses

// ApiResponse A generic API response struct. Typing should be stricter, but this is a small app and
// functionally a demo
type ApiResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
