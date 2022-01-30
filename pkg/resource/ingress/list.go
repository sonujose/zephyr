package ingress

import (
	v1 "k8s.io/api/networking/v1"
)

func (r *resource) ListIngress(namespace *string, selectors *map[string]string, clusterScope bool) (*v1.IngressList, error) {

	ingressListChannel := GetIngressListChannel(r.kclient, namespace, clusterScope)

	ingress := <-ingressListChannel.List
	err := <-ingressListChannel.Error

	return ingress, err
}
