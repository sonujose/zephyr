package resource

import (
	v1 "k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListServices(namespace *string) (*v1.ServiceList, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
