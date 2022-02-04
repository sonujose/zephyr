package service

import (
	"encoding/json"

	"github.com/sonujose/kube-spectrum/pkg/resource/ingress"
)

type IngressServiceMap struct {
	Metadata ingress.Metadata `json:"metadata"`
	Routes   ingress.Paths    `json:"spec"`
	Status   ingress.Status   `json:"status"`
}

/*
	ListServiceMappings
	 - List all ingress in the namespace
	 - Get details of the input service mostly selectors
*/
func (r *resource) ListServiceMappings(namespace *string, service *string) (*[]IngressServiceMap, error) {

	ingressListChannel := ingress.GetIngressListChannel(r.kclient, namespace, false)

	ingresslist := <-ingressListChannel.List
	err := <-ingressListChannel.Error

	if err != nil {
		return nil, err
	}

	var ingressItem []ingress.IngressDto

	b, _ := json.Marshal(ingresslist.Items)
	err = json.Unmarshal(b, &ingressItem)

	var rs []IngressServiceMap

	for _, j := range ingressItem {
		for k := 0; k < len(j.Spec.Rules); k++ {
			for _, path := range j.Spec.Rules[k].HTTP.Paths {
				if path.Backend.Service.Name == *service {
					svcmap := &IngressServiceMap{Status: j.Status, Metadata: j.Metadata, Routes: path}
					rs = append(rs, *svcmap)
				}
			}
		}

	}

	return &rs, err
}
