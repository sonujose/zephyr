package kube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/api/dto"
	"github.com/sonujose/kube-spectrum/internal/logger"
	"github.com/sonujose/kube-spectrum/pkg/resource/namespace"
)

// GetNamespaces godoc
// @Summary Get list of namespace
// @Tags Namespace
// @Accept */*
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []dto.NamespaceListResponse
// @Failure 500 {object} []dto.ErrorResponse
// @Router /namespaces [get]
func (h *apihandler) GetNamespaces(c *gin.Context) {

	// Get the context aware logger set by the Logging middleware
	logmanager := logger.GetContextAwareLogger(c)

	resource := namespace.New(KubeClient)

	ns, err := resource.ListNamespaces()

	if err != nil {

		logmanager.WithFields(logrus.Fields{"error": err}).Errorf("Error fetching namespace list from cluster")

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:    "Operation failed",
			IsSuccess: false,
			Error:     "Internal Server Error, Unable to fetch list of namespaces",
		})

		return
	}

	logmanager.Debugf("Successfully fetched %d namespaces", len(*ns))

	c.JSON(http.StatusOK, &dto.NamespaceListResponse{
		IsSuccess: true,
		Data:      ns,
	})
}
