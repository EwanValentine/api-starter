package handlers

// Response - Http response
type Response struct {
	Data interface{}            `json:"data"`
	Meta map[string]interface{} `json:"_meta"`
}
