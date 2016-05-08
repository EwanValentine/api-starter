package handlers

import "net/http"

// Error - Http error
type Error struct {
	Message string `json:"_message"`
	Code    int    `json:"_code"`
}

var (
	// NotFound - Basic 404 response
	NotFound = &Error{
		Message: "Resource not found",
		Code:    http.StatusNotFound,
	}

	// Unprocessable - 422 response type
	Unprocessable = &Error{
		Message: "Entity unprocessable",
		Code:    422,
	}
)
