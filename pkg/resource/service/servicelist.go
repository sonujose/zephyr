package service

import (
	"github.com/sonujose/kube-spectrum/pkg/resource/pod"
)

func (r *resource) ListServices(namespace *string) (*[]ServiceDto, error) {
	servicechannel := GetServiceListChannel(r.kclient, namespace)

	services := <-servicechannel.List
	err := <-servicechannel.Error

	if err != nil {
		return nil, err
	}

	var svclist []ServiceDto

	resource := pod.New(r.kclient)

	// Creating service object for unerstanding pod status
	for _, k := range services.Items {
		serviceItem := &ServiceDto{
			Name:              k.Name,
			Namespace:         k.Namespace,
			Labels:            k.Labels,
			Type:              string(k.Spec.Type),
			Annotations:       k.Annotations,
			CreationTimestamp: k.CreationTimestamp.Time,
			Selector:          k.Spec.Selector,
			ClusterIP:         k.Spec.ClusterIP,
			ExternalIP:        k.Spec.ExternalIPs[0],
		}

		// Setting default state to be failed, Success state is set only if one of pod is in ready state
		serviceItem.State = "Failed"

		// Fetching all pods underlying the container
		pods, err := resource.ListPods(namespace, &k.Spec.Selector)

		if err != nil {
			svclist = append(svclist, *serviceItem)
			continue
		}

		var podsList []PodInfo

		// Understanding the pod ready status
		// If any one container is failing pod isReady is false
		for _, pod := range pods.Items {
			podItem := &PodInfo{
				Name:   pod.Name,
				Status: string(pod.Status.Phase),
				Reason: pod.Status.Reason,
			}

			// Optimistic - container is always ready..
			// container should run!! thats the purpose, its strange if it is failing. Your views please..
			containerstatus := true

			for _, cs := range pod.Status.ContainerStatuses {
				if !cs.Ready {
					containerstatus = false
				}
			}

			// If all the containers are in ready state, pod status is ready
			// Since one of the pod status is ready service state will be success
			if containerstatus {
				serviceItem.State = "Active"
				podItem.IsReady = true
			}

			podsList = append(podsList, *podItem)
		}

		serviceItem.Pods = podsList
		svclist = append(svclist, *serviceItem)
	}

	return &svclist, err
}
