package service

import (
	"encoding/json"

	"github.com/sonujose/kube-spectrum/pkg/resource/ingress"
)

type IngressServiceMap struct {
	Metadata     ingress.Metadata `json:"metadata"`
	Routes       ingress.Paths    `json:"spec"`
	Status       ingress.Status   `json:"status"`
	IngressClass string           `json:"ingressClass"`
}

/*
	ListServiceMappings
	 - Maps all ingress defined in the namespace with the given service
	 - Get details of the input service mostly selectors
	 - It traverses throgh all possible ways ingress can be defined in the control
	 - Even if same service is defines under multiple host or multiple path, the traverser will find it
	 - It also check if service is upstreamed via multiple ingress objects
*/
func (r *resource) ListServiceMappings(namespace *string, service *string) (*[]IngressServiceMap, error) {

	ingressListChannel := ingress.GetIngressListChannel(r.kclient, namespace, false)

	ingresslist := <-ingressListChannel.List
	err := <-ingressListChannel.Error

	if err != nil {
		return nil, err
	}

	// Converting v1.Ingress to IngressDto to map service metdatda
	// along with the ingress rules
	var ingressListDto []ingress.IngressDto
	ingressListObject, _ := json.Marshal(ingresslist.Items)
	err = json.Unmarshal(ingressListObject, &ingressListDto)

	if err != nil {
		return nil, err
	}

	// Lopping through all ingress ites, multiple rules and multiple paths
	// to find the ingress paths referred to the specified service
	// Traverse through multi host object as well as different path under same host or even different ingress object
	var ingressSvcMap []IngressServiceMap
	for _, j := range ingressListDto { // Loop through all the ingress objects
		for k := 0; k < len(j.Spec.Rules); k++ { // Loop through all the ingress rules
			for _, path := range j.Spec.Rules[k].HTTP.Paths { // Loop through all the ingress paths
				// Check if the backend service is defined as the specified one
				// If found true, create a servicemap object with all metdata and current status
				if path.Backend.Service.Name == *service {
					ingressclass := j.Metadata.Annotations["kubernetes.io/ingress.class"]
					svcmap := &IngressServiceMap{Status: j.Status, Metadata: j.Metadata, Routes: path, IngressClass: ingressclass}
					ingressSvcMap = append(ingressSvcMap, *svcmap)
				}
			}
		}

	}

	return &ingressSvcMap, err
}
