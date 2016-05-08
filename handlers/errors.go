package handlers

// Error - Http error
type Error struct {
	Message string `json:"_message"`
	Code    int    `json:"_code"`
}
