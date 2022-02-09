package pod

import (
	v1 "k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListPods(namespace *string, selectors *map[string]string) (*v1.PodList, error)
	ListPodsDetailByService(namespace *string, service *string) (*PodDto, error)
	GetPodLogsBlob(namespace *string, pod *string, container *string) (*string, error)
}

type resource struct {
	kclient *client.Clientset
}

func New(client *client.Clientset) Resource {
	return &resource{kclient: client}
}
