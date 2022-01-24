package namespace

import (
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListNamespaces() (*[]string, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
