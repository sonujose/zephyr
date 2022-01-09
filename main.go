package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	apis "github.com/sonujose/kube-spectrum/api"
	"github.com/sonujose/kube-spectrum/api/handler/kube"
	"github.com/sonujose/kube-spectrum/internal/logger"
	"github.com/sonujose/kube-spectrum/internal/utility"
)

// @title Kube-spectrum kubernetes Service
// @version 2.0
// @description Kube-spectrum kubernetes Dashboard Service
// @team Kube-spectrum

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /api/v1
func main() {

	log := logger.Initialize()
	log.Debug("Starting kube spectrum server")

	// Initializing gin gonic server in release mode
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	// Middleware to setup Logmanager
	// SetCorrelationID - middleware is fired for all api requests initialing new logrus for handler
	// RequestLogger - middleware logging the request info for the API request
	router.Use(logger.SettraceID(log))
	router.Use(logger.RequestLogger(log))

	// Gin cors policy and recovery mode for failures
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	// Calling routes method to register all api routes
	apis.RegisterAPIRoutes(router, log)

	kube.InitClient()

	log.Infof("Server listening on port %s", utility.GetEnv("APP_PORT", "7500"))

	router.Run(":" + utility.GetEnv("APP_PORT", "7500"))
}
