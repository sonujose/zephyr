package dto

type ServiceResultResponse struct {
	IsSuccess bool        `json:"isSuccess"`
	Data      interface{} `json:"data"`
}
