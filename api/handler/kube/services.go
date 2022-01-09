package kube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonujose/kube-spectrum/api/dto"
	"github.com/sonujose/kube-spectrum/internal/logger"
)

// GetServices godoc
// @Summary Get services for specified namespace
// @Tags Services
// @Param namespace  path	string	true	"Namespace"
// @Accept */*
// @Produce json
// @Success 200 {object} []dto.ServiceResultResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /api/services/{namespace} [get]
func (h *apihandler) GetServices(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	namespace := c.Param("namespace")

	//kclient := resources.NewClient()

	logmanager.Infof("Successfully fetched Services for the namespace %s", namespace)

	c.JSON(http.StatusOK, &dto.ServiceResultResponse{
		IsSuccess: true,
		Data:      nil,
	})
}
