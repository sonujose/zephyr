package pod

import "time"

type PodDto struct {
	// Name of the pod
	Name string `json:"name"`

	// Namespace where pod is created
	Namespace string `json:"namespace"`

	// Time when the pod was created
	CreationTimestamp time.Time `json:"creationTimestamp"`

	// List of all labels associated with the pod
	Labels map[string]string `json:"labels"`

	// List of all annotations associated with the pod
	Annotations map[string]string `json:"annotations"`
}
