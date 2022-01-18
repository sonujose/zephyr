package resource

import (
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListServices(namespace *string) (*[]ServiceDto, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
