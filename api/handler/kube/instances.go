package kube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/api/dto"
	"github.com/sonujose/kube-spectrum/internal/logger"
	"github.com/sonujose/kube-spectrum/pkg/resource/pod"
)

// GetServiceDetails godoc
// @Summary Get full details of the service in the specified namespace
// @Tags Services
// @Param namespace  path	string	true "namespace"
// @Param service  path	string	true "service"
// @Accept */*
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []dto.ServiceDetailsResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /services/{namespace}/{service} [get]
func (h *apihandler) GetServiceDetails(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	namespace := c.Param("namespace")
	serviceName := c.Param("service")

	if len(serviceName) == 0 || len(namespace) == 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:    "Operation failed, Provide service name and namespace",
			IsSuccess: false,
			Error:     "Bad request - Misssing fields",
		})

		return
	}

	resource := pod.New(KubeClient)

	service, err := resource.ListPodsDetailByService(&namespace, &serviceName)

	if err != nil {

		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Error fetching service details of %s from cluster for namespace %s", serviceName, namespace)

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:    "Operation failed",
			IsSuccess: false,
			Error:     "Internal Server Error, Unable to fetch service details from cluster",
		})

		return
	}

	c.JSON(http.StatusOK, &dto.ServiceDetailsResponse{
		IsSuccess: true,
		Message:   *service,
	})
}
