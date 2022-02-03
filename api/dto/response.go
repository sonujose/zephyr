package dto

import (
	"github.com/sonujose/kube-spectrum/pkg/resource/ingress"
	"github.com/sonujose/kube-spectrum/pkg/resource/service"
)

type ServiceResultResponse struct {
	IsSuccess bool                 `json:"isSuccess"`
	Data      []service.ServiceDto `json:"message"`
}

type NamespaceListResponse struct {
	IsSuccess bool      `json:"isSuccess"`
	Data      *[]string `json:"message"`
}

type ServiceDetailsResponse struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   interface{} `json:"message"`
}

type IngressDetailsResponse struct {
	IsSuccess bool                  `json:"isSuccess"`
	Message   *[]ingress.IngressDto `json:"message"`
}
