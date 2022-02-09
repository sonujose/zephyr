package pod

import (
	"context"
	"io"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
)

type PodsListChannel struct {
	List  chan *v1.PodList
	Error chan error
}

/**
  GetPodsListChannel
  Routine return pods channel to fetch pods based input labels
**/
func GetPodsListChannel(client client.Interface, namespace *string, labelselector *map[string]string) PodsListChannel {

	d := &metav1.LabelSelector{
		MatchLabels: *labelselector,
	}

	selector, _ := metav1.LabelSelectorAsSelector(d)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	channelSize := 1

	channel := PodsListChannel{
		List:  make(chan *v1.PodList, 1),
		Error: make(chan error, 1),
	}

	go func() {
		list, err := client.CoreV1().Pods(*namespace).List(context.TODO(), options)
		for i := 0; i < channelSize; i++ {
			channel.List <- list
			channel.Error <- err
		}
	}()

	return channel
}

/**
  GetPodsLogsChannel
  Routine return pods channel to fetch logs based input
**/
func GetPodsLogsChannel(client client.Interface, namespace *string, pod *string, container *string) (*string, error) {

	podLogOptions := v1.PodLogOptions{
		Container: *container,
		Follow:    true,
	}

	podLogRequest := client.CoreV1().Pods(*namespace).GetLogs(*pod, &podLogOptions)

	stream, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return nil, err
	}

	defer stream.Close()

	var logBlobData string
	for {
		buf := make([]byte, 2000)
		numBytes, err := stream.Read(buf)
		if numBytes == 0 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		logBlobData = string(buf[:numBytes])
		// fmt.Print(logBlobData)
	}

	return &logBlobData, nil
}
