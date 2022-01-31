package pod

import "time"

type PodDto struct {
	// Name of the pod
	Service string `json:"service"`

	// Namespace where pod is created
	Namespace string `json:"namespace"`

	// Spec and metadata related to PodInfo
	Info *[]PodInfo `json:"info"`
}

type PodInfo struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Metadata  Metadata `json:"metadata"`
	Spec      Spec     `json:"spec"`
	Status    Status   `json:"status"`
}
type OwnerReferences struct {
	APIVersion         string `json:"apiVersion"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	UID                string `json:"uid"`
	Controller         bool   `json:"controller"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
}
type Metadata struct {
	Name              string            `json:"name"`
	GenerateName      string            `json:"generateName"`
	Namespace         string            `json:"namespace"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	ResourceVersion   string            `json:"resourceVersion"`
	CreationTimestamp time.Time         `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	OwnerReferences   []OwnerReferences `json:"ownerReferences"`
}
type Secret struct {
	SecretName  string `json:"secretName"`
	DefaultMode int    `json:"defaultMode"`
}
type Volumes struct {
	Name   string `json:"name"`
	Secret Secret `json:"secret"`
}
type Ports struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}
type ConfigMapRef struct {
	Name string `json:"name"`
}
type EnvFrom struct {
	ConfigMapRef ConfigMapRef `json:"configMapRef"`
}
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Limits struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type Requests struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type Resources struct {
	Limits   Limits   `json:"limits"`
	Requests Requests `json:"requests"`
}
type VolumeMounts struct {
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readOnly"`
	MountPath string `json:"mountPath"`
}
type SecurityContext struct {
	RunAsUser  int `json:"runAsUser"`
	RunAsGroup int `json:"runAsGroup"`
}
type Containers struct {
	Name                     string          `json:"name"`
	Image                    string          `json:"image"`
	Ports                    []Ports         `json:"ports"`
	EnvFrom                  []EnvFrom       `json:"envFrom"`
	Env                      []Env           `json:"env"`
	Resources                Resources       `json:"resources"`
	VolumeMounts             []VolumeMounts  `json:"volumeMounts"`
	TerminationMessagePath   string          `json:"terminationMessagePath"`
	TerminationMessagePolicy string          `json:"terminationMessagePolicy"`
	ImagePullPolicy          string          `json:"imagePullPolicy"`
	SecurityContext          SecurityContext `json:"securityContext"`
}

type ImagePullSecrets struct {
	Name string `json:"name"`
}
type Tolerations struct {
	Key               string `json:"key"`
	Operator          string `json:"operator"`
	Effect            string `json:"effect"`
	TolerationSeconds int    `json:"tolerationSeconds,omitempty"`
}
type Spec struct {
	Volumes                       []Volumes          `json:"volumes"`
	Containers                    []Containers       `json:"containers"`
	RestartPolicy                 string             `json:"restartPolicy"`
	TerminationGracePeriodSeconds int                `json:"terminationGracePeriodSeconds"`
	DNSPolicy                     string             `json:"dnsPolicy"`
	ServiceAccountName            string             `json:"serviceAccountName"`
	ServiceAccount                string             `json:"serviceAccount"`
	NodeName                      string             `json:"nodeName"`
	SecurityContext               SecurityContext    `json:"securityContext"`
	ImagePullSecrets              []ImagePullSecrets `json:"imagePullSecrets"`
	SchedulerName                 string             `json:"schedulerName"`
	Tolerations                   []Tolerations      `json:"tolerations"`
	Priority                      int                `json:"priority"`
	EnableServiceLinks            bool               `json:"enableServiceLinks"`
	PreemptionPolicy              string             `json:"preemptionPolicy"`
}
type Conditions struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastProbeTime      interface{} `json:"lastProbeTime"`
	LastTransitionTime time.Time   `json:"lastTransitionTime"`
}
type PodIPs struct {
	IP string `json:"ip"`
}
type Running struct {
	StartedAt time.Time `json:"startedAt"`
}
type State struct {
	Running Running `json:"running"`
}
type LastState struct {
}
type ContainerStatuses struct {
	Name         string    `json:"name"`
	State        State     `json:"state"`
	LastState    LastState `json:"lastState"`
	Ready        bool      `json:"ready"`
	RestartCount int       `json:"restartCount"`
	Image        string    `json:"image"`
	ImageID      string    `json:"imageID"`
	ContainerID  string    `json:"containerID"`
	Started      bool      `json:"started"`
}
type Status struct {
	Phase             string              `json:"phase"`
	Conditions        []Conditions        `json:"conditions"`
	HostIP            string              `json:"hostIP"`
	PodIP             string              `json:"podIP"`
	PodIPs            []PodIPs            `json:"podIPs"`
	StartTime         time.Time           `json:"startTime"`
	ContainerStatuses []ContainerStatuses `json:"containerStatuses"`
	QosClass          string              `json:"qosClass"`
}
