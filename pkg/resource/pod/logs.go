package pod

func (r *resource) GetPodLogsBlob(namespace *string, pod *string, container *string) (*string, error) {

	podLogsBlob, err := GetPodsLogsChannel(r.kclient, namespace, pod, container)

	return podLogsBlob, err
}
