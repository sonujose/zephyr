package dto

// ErrorResponse
type ErrorResponse struct {
	Status    string `json:"status"`
	Error     string `json:"error"`
	IsSuccess bool   `json:"isSuccess"`
}
