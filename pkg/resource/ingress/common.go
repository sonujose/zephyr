package ingress

import (
	v1 "k8s.io/api/networking/v1"
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListIngress(namespace *string, selectors *map[string]string, clusterScope bool) (*v1.IngressList, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
