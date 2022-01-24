package dto

import (
	"github.com/sonujose/kube-spectrum/pkg/resource/service"
)

type ServiceResultResponse struct {
	IsSuccess bool                 `json:"isSuccess"`
	Data      []service.ServiceDto `json:"data"`
}

type NamespaceListResponse struct {
	IsSuccess bool      `json:"isSuccess"`
	Data      *[]string `json:"data"`
}
