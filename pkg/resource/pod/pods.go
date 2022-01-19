package pod

import (
	v1 "k8s.io/api/core/v1"
)

func (r *resource) ListPods(namespace *string, selectors *map[string]string) (*v1.PodList, error) {

	podsListChannel := GetPodsListChannel(r.kclient, namespace, selectors)

	pods := <-podsListChannel.List
	err := <-podsListChannel.Error

	return pods, err
}
