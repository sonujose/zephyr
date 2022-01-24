package namespace

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
	ListNamespaces
	List all namespace in the cluster - return only the name
**/
func (r *resource) ListNamespaces() (*[]string, error) {
	list, err := r.kclient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var namespaceList []string
	for _, ns := range list.Items {
		namespaceList = append(namespaceList, ns.Name)
	}

	return &namespaceList, err
}
