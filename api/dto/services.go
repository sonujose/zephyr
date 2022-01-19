package dto

import "github.com/sonujose/kube-spectrum/pkg/resource"

type ServiceResultResponse struct {
	IsSuccess bool                  `json:"isSuccess"`
	Data      []resource.ServiceDto `json:"data"`
}
