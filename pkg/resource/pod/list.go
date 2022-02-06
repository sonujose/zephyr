package pod

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/internal/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *resource) ListPods(namespace *string, selectors *map[string]string) (*v1.PodList, error) {

	podsListChannel := GetPodsListChannel(r.kclient, namespace, selectors)

	pods := <-podsListChannel.List
	err := <-podsListChannel.Error

	return pods, err
}

func (r *resource) ListPodsDetailByService(namespace *string, service *string) (*PodDto, error) {

	logmanager := logger.Get()
	servicedata, err := r.kclient.CoreV1().Services(*namespace).Get(context.TODO(), *service, metav1.GetOptions{})

	if err != nil {
		logmanager.WithFields(logrus.Fields{"error": err}).Warnf("Unable to fetch details Provided service %s in the namespace %s", *service, *namespace)
		return nil, err
	}

	ServiceSelectors := servicedata.Spec.Selector

	podsListChannel := GetPodsListChannel(r.kclient, namespace, &ServiceSelectors)

	pods := <-podsListChannel.List
	err = <-podsListChannel.Error

	if err != nil {
		return nil, err
	}

	var podInfo []PodInfo

	b, _ := json.Marshal(pods.Items)
	err = json.Unmarshal(b, &podInfo)

	if err != nil {
		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Failed to parse Pod information from the cluster for service %s", *service)
		return nil, err
	}

	return &PodDto{Service: *service, Info: &podInfo, Namespace: *namespace, Selectors: &ServiceSelectors}, err
}
