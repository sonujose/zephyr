package ingress

import (
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListIngress(namespace *string, clusterScope bool) (*[]IngressDto, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
