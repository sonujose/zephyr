package resource

import (
	v1 "k8s.io/api/core/v1"
)

func (r *resource) ListServices(namespace *string) (*[]ServiceDto, error) {
	servicechannel := GetServiceListChannel(r.kclient, namespace)

	services := <-servicechannel.List
	err := <-servicechannel.Error

	return toServiceDTO(services), err
}

func toServiceDTO(servicesList *v1.ServiceList) *[]ServiceDto {
	var svclist []ServiceDto

	for _, k := range servicesList.Items {
		serviceItem := &ServiceDto{
			Name:              k.Name,
			Namespace:         k.Namespace,
			Labels:            k.Labels,
			Type:              string(k.Spec.Type),
			Annotations:       k.Annotations,
			CreationTimestamp: k.CreationTimestamp.Time,
			Selector:          k.Spec.Selector,
		}

		svclist = append(svclist, *serviceItem)
	}

	return &svclist
}
