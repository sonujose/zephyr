package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sonujose/kube-spectrum/internal/logger"
	kclient "github.com/sonujose/kube-spectrum/pkg/client"
	"k8s.io/client-go/kubernetes"
)

type APIHandler interface {
	GetServices(c *gin.Context)
	GetNamespaces(c *gin.Context)
	GetServiceDetails(c *gin.Context)
	GetIngress(c *gin.Context)
}

type apihandler struct {
	kclient *kubernetes.Clientset
}

func NewHandler() APIHandler {
	return &apihandler{kclient: KubeClient}
}

var (
	KubeClient *kubernetes.Clientset
)

func InitClient() {
	log := logger.Get()
	client, err := kclient.NewKubeClient()

	if err != nil {
		log.Errorf("Error initialing kubernetes client")
		return
	}

	KubeClient = client

	log.Infof("Successfully initialized kubernetes client.")
}
