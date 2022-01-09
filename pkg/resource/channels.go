package resource

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
)

type ServiceListChannel struct {
	List  chan *v1.ServiceList
	Error chan error
}

func GetServiceListChannel(client client.Interface, namespace *string) ServiceListChannel {

	channelSize := 1

	channel := ServiceListChannel{
		List:  make(chan *v1.ServiceList, 1),
		Error: make(chan error, 1),
	}

	go func() {
		list, err := client.CoreV1().Services(*namespace).List(context.TODO(), metav1.ListOptions{})
		for i := 0; i < channelSize; i++ {
			channel.List <- list
			channel.Error <- err
		}
	}()

	return channel
}
