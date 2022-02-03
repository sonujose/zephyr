package ingress

import "time"

type IngressDto struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}
type Metadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	ResourceVersion   string            `json:"resourceVersion"`
	Generation        int               `json:"generation"`
	CreationTimestamp time.Time         `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}
type Port struct {
	Number int `json:"number"`
}
type Service struct {
	Name string `json:"name"`
	Port Port   `json:"port"`
}
type Backend struct {
	Service Service `json:"service"`
}
type Paths struct {
	Path     string  `json:"path"`
	PathType string  `json:"pathType"`
	Backend  Backend `json:"backend"`
}
type HTTP struct {
	Paths []Paths `json:"paths"`
}
type Rules struct {
	Host string `json:"host"`
	HTTP HTTP   `json:"http"`
}
type Spec struct {
	Rules []Rules `json:"rules"`
}
type Ingress struct {
	IP string `json:"ip"`
}
type LoadBalancer struct {
	Ingress []Ingress `json:"ingress"`
}
type Status struct {
	LoadBalancer LoadBalancer `json:"loadBalancer"`
}
