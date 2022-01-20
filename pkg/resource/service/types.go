package service

import (
	"time"
)

type ServiceDto struct {
	// Name of the service
	Name string `json:"name"`

	// Namespace where service is created
	Namespace string `json:"namespace"`

	// Time when the service was created
	CreationTimestamp time.Time `json:"creationTimestamp"`

	// List of all labels associated with the service
	Labels map[string]string `json:"labels"`

	// List of all annotations associated with the service
	Annotations map[string]string `json:"annotations"`

	// Label selector of the service.
	Selector map[string]string `json:"selector"`

	// ClusterIP is usually assigned by the master. Valid values are None, empty string (""), or
	// a valid IP address. None can be specified for headless services when proxying is not required
	ClusterIP string `json:"clusterIP"`

	// ExternalIP is the Ip of the Loadbalancer attached with the service
	// If service type is not Loadbalaner then IP will be none
	ExternalIP string `json:"externalIP"`

	// Type determines how the service will be exposed.  Valid options: ClusterIP, NodePort, LoadBalancer, ExternalName
	Type string `json:"type"`

	// Ports mapped to the service
	Ports []Ports `json:"ports"`

	// Pod Instances
	Pods []PodInfo `json:"podInfo"`

	// Success statsus, check if any one pod is available to pass request
	State string `json:"state"`
}

type ServiceType string

type PodInfo struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
	IsReady bool   `json:"isReady"`
}

type Ports struct {
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
	Port       int    `json:"port"`
	TargetPort string `json:"targetPort"`
}
