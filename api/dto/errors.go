package dto

// ErrorResponse
type ErrorResponse struct {
	Status    string `json:"description"`
	Error     string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}
