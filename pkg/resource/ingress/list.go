package ingress

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/internal/logger"
)

func (r *resource) ListIngress(namespace *string, clusterScope bool) (*[]IngressDto, error) {

	logmanager := logger.Get()
	ingressListChannel := GetIngressListChannel(r.kclient, namespace, clusterScope)

	ingress := <-ingressListChannel.List
	err := <-ingressListChannel.Error

	if err != nil {
		return nil, err
	}

	var ingressItem []IngressDto

	b, _ := json.Marshal(ingress.Items)
	err = json.Unmarshal(b, &ingressItem)

	if err != nil {
		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Failed to parse ingress information from the cluster for namespace", *namespace)
		return nil, err
	}

	return &ingressItem, err
}
