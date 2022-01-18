package resource

import (
	"context"

	"github.com/sonujose/kube-spectrum/internal/utility"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	client "k8s.io/client-go/kubernetes"
)

type ServiceListChannel struct {
	List  chan *v1.ServiceList
	Error chan error
}

/**
  GetLabelSelectorListOptionsForService
  Fetch the global label seletor key value pair form the enviornment
  Supply the LabelSelector to service List options
**/
func GetLabelSelectorListOptionsForService() metav1.ListOptions {
	labelkey := utility.GetEnv("SERVICE_LABEL_SELECTOR_KEY", "")
	labelValue := utility.GetEnv("SERVICE_LABEL_SELECTOR_VALUE", "")

	var selector labels.Selector

	if labelkey != "" && labelValue != "" {
		d := &metav1.LabelSelector{
			MatchLabels: map[string]string{labelkey: labelValue},
		}
		selector, _ = metav1.LabelSelectorAsSelector(d)

		return metav1.ListOptions{LabelSelector: selector.String()}
	}

	return metav1.ListOptions{}
}

// Routine to get create channel to fetch services based on configured labels
func GetServiceListChannel(client client.Interface, namespace *string) ServiceListChannel {

	channelSize := 1

	channel := ServiceListChannel{
		List:  make(chan *v1.ServiceList, 1),
		Error: make(chan error, 1),
	}

	serviceListOptions := GetLabelSelectorListOptionsForService()
	go func() {
		list, err := client.CoreV1().Services(*namespace).List(context.TODO(), serviceListOptions)
		for i := 0; i < channelSize; i++ {
			channel.List <- list
			channel.Error <- err
		}
	}()

	return channel
}
