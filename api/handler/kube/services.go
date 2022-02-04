package kube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/api/dto"
	"github.com/sonujose/kube-spectrum/internal/logger"
	"github.com/sonujose/kube-spectrum/pkg/resource/service"
)

// GetServices godoc
// @Summary Get services for specified namespace
// @Tags Services
// @Param namespace  path	string	false "namespace"
// @Accept */*
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []dto.ServiceResultResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /services/{namespace} [get]
func (h *apihandler) GetServices(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	namespace := c.Param("namespace")

	clusterScope := false
	if len(namespace) == 0 || namespace == "all-ns" {
		clusterScope = true
	}

	resource := service.New(KubeClient)

	services, err := resource.ListServices(&namespace, clusterScope)

	if err != nil {

		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Error fetching services from cluster for namespace %s", namespace)

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:    "Operation failed",
			IsSuccess: false,
			Error:     "Internal Server Error, Unable to fetch services from cluster",
		})

		return
	}

	logmanager.Debugf("Successfully fetched %d Services for the namespace %s", len(*services), namespace)

	c.JSON(http.StatusOK, &dto.ServiceResultResponse{
		IsSuccess: true,
		Data:      *services,
	})
}

// GetServiceMappingWithIngress godoc
// @Summary Get ingress mapping details for specified service
// @Tags Services
// @Param service  path	string	true "service"
// @Param namespace  path	string	true "namespace"
// @Accept */*
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []dto.ServiceResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /services/mappings/ingress/{namespace}/{service} [get]
func (h *apihandler) GetServiceMappingWithIngress(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	namespace := c.Param("namespace")
	serviceName := c.Param("service")

	if len(namespace) == 0 || len(serviceName) == 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:    "Bad Request",
			IsSuccess: false,
			Error:     "Bad request, specify namespace and service name",
		})

		return
	}

	resource := service.New(KubeClient)

	serviceMappings, err := resource.ListServiceMappings(&namespace, &serviceName)

	if err != nil {

		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Error fetching ingress mappings for the service %s for namespace %s", serviceName, namespace)

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:    "Operation failed",
			IsSuccess: false,
			Error:     "Internal Server Error, Unable to fetch ingress mappings for service",
		})

		return
	}

	c.JSON(http.StatusOK, &dto.ServiceResponse{
		IsSuccess: true,
		Message:   *serviceMappings,
	})
}
