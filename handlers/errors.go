package handlers

import "net/http"

// Error - Http error
type Error struct {
	Message string `json:"_message"`
	Code    int    `json:"_code"`
}

var (
	NotFound = &Error{
		Message: "Resource not found",
		Code:    http.StatusNotFound,
	}

	Unprocessable = &Error{
		Message: "Entity unprocessable",
		Code:    422,
	}
)
