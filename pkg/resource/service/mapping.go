package service

import (
	"fmt"

	"github.com/sonujose/kube-spectrum/pkg/resource/ingress"
)

type IngressServiceMap struct {
	Metadata ingress.Metadata `json:"metadata"`
	Routes   []ingress.Paths  `json:"spec"`
	Status   ingress.Status   `json:"status"`
}

/*
	TODO:
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

	fmt.Println(ingresslist)

	var rs []IngressServiceMap

	// for i, ingressObj := range ingresslist.Items {

	// 	if j.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Name == *service {
	// 		ingressObj.me
	// 	}
	// }

	return &rs, err
}
