package api

// ErrorResponse is a struct that contains error message.
// This should be used for all error responses in http server.
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateErrorResponse creates ErrorResponse from string.
func CreateErrorResponse(err string) *ErrorResponse {
	return &ErrorResponse{
		Error: err,
	}
}
