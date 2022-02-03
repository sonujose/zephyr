package kube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/api/dto"
	"github.com/sonujose/kube-spectrum/internal/logger"
	"github.com/sonujose/kube-spectrum/pkg/resource/ingress"
)

// GetIngress godoc
// @Summary Get services for specified namespace
// @Tags Services
// @Param namespace  path	string	false "namespace"
// @Accept */*
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []dto.IngressDetailsResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /ingress/{namespace} [get]
func (h *apihandler) GetIngress(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	namespace := c.Param("namespace")

	clusterScope := false
	if len(namespace) == 0 || namespace == "all-ns" {
		clusterScope = true
	}

	resource := ingress.New(KubeClient)

	ingressList, err := resource.ListIngress(&namespace, clusterScope)

	if err != nil {

		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Error fetching ingress from cluster for namespace %s", namespace)

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:    "Operation failed",
			IsSuccess: false,
			Error:     "Internal Server Error, Unable to fetch ingress from cluster",
		})

		return
	}

	logmanager.Debugf("Successfully fetched %d ingress for the namespace %s", len(*ingressList), namespace)

	c.JSON(http.StatusOK, &dto.IngressDetailsResponse{
		IsSuccess: true,
		Message:   ingressList,
	})
}
