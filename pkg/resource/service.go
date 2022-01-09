package resource

import (
	v1 "k8s.io/api/core/v1"
)

func (r *resource) ListServices(namespace *string) (*v1.ServiceList, error) {
	servicechannel := GetServiceListChannel(r.kclient, namespace)

	services := <-servicechannel.List
	err := <-servicechannel.Error

	return services, err
}
