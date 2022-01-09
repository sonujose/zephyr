package kube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/api/dto"
	"github.com/sonujose/kube-spectrum/internal/logger"
	kresource "github.com/sonujose/kube-spectrum/pkg/resource"
)

// GetServices godoc
// @Summary Get services for specified namespace
// @Tags Services
// @Param namespace  path	string	true	"Namespace"
// @Accept */*
// @Produce json
// @Success 200 {object} []dto.ServiceResultResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /services/{namespace} [get]
func (h *apihandler) GetServices(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	namespace := c.Param("namespace")

	res := kresource.New(KubeClient)

	services, err := res.ListServices(&namespace)

	if err != nil {

		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Error fetching services from cluster for namespace %s", namespace)

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:    "Operation failed",
			IsSuccess: false,
			Error:     "Internal Server Error, Unable to fetch services from cluster",
		})

		return
	}

	logmanager.Infof("Successfully fetched Services for the namespace %s", namespace)

	c.JSON(http.StatusOK, &dto.ServiceResultResponse{
		IsSuccess: true,
		Data:      *services,
	})
}
