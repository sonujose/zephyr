package service

import (
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListServices(namespace *string, clusterScopr bool) (*[]ServiceDto, error)
	ListServiceMappings(namespace *string, service *string) (*[]IngressServiceMap, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
