package kube

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

type APIHandler interface {
	GetConfigmapList(c *gin.Context)
}

type apihandler struct{}

func New() APIHandler {
	return &apihandler{}
}

var (
	KubeClient resources.KubeClient
)

func InitKubeClient() {
	KubeClient, err := resources.NewClient()

	if err != nil {
		klog.Errorf("Error initialing kubernetes client")

	}

	klog.Infof("Successfully initialized kubernetes client %v", KubeClient)
}
