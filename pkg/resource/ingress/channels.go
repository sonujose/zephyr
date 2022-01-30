package ingress

import (
	"context"

	"github.com/sonujose/kube-spectrum/internal/utility"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	client "k8s.io/client-go/kubernetes"
)

type ServiceListChannel struct {
	List  chan *v1.IngressList
	Error chan error
}

/**
  GetServiceListChannel
  Routine return servicelist channel to fetch services based on configured labels
**/
func GetIngressListChannel(client client.Interface, namespace *string, clusterScope bool) ServiceListChannel {

	channelSize := 1

	channel := ServiceListChannel{
		List:  make(chan *v1.IngressList, 1),
		Error: make(chan error, 1),
	}

	// Setting namespace to empty if needed to list services from all namespaces
	var allnamespaceQuery = ""
	if clusterScope {
		*namespace = allnamespaceQuery
	}

	serviceListOptions := GetLabelSelectorListOptionsForService()
	go func() {
		list, err := client.NetworkingV1().Ingresses(*namespace).List(context.TODO(), serviceListOptions)

		for i := 0; i < channelSize; i++ {
			channel.List <- list
			channel.Error <- err
		}
	}()

	return channel
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
