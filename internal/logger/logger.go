package logger

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sonujose/kube-spectrum/internal/utility"
)

// CorrelationHeader defines a default Correlation ID HTTP header.
const (
	CorrelationHeader  = "x-correlation-id"
	ContextKey         = "traceID"
	Contextawarelogger = "contextawarelogger"
	LogTimestampFormat = "2006-01-02 15:04:05"
)

var (
	Logger *logrus.Entry
)

// Log manager Initializer
// Initialize log module using logrus
func Initialize() *logrus.Logger {
	log := logrus.New()

	// Setting up format for logrus module
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: LogTimestampFormat,
		FullTimestamp: true, ForceColors: true, DisableLevelTruncation: true})

	log.SetLevel(getLowestLoggingLevel())
	log.SetOutput(logrus.StandardLogger().Out)

	return log
}

// SetTraceId (HTTP Middleware)
// Initialize Logmanager instance
// Handler to set Logrus with custom context
func SettraceID(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Checking for traceID in the request header, if not found will create a new uuid
		traceID := c.Request.Header.Get(CorrelationHeader)
		if traceID == "" {
			traceID = uuid.New().String()
		}

		//path := c.Request.URL.Path
		// Adding correlation-id in response headers for every request
		c.Writer.Header().Set("x-correlation-id", traceID)

		logmanager := logger.WithFields(logrus.Fields{ContextKey: traceID})

		c.Set(Contextawarelogger, logmanager)
		c.Set(ContextKey, traceID)

		// if !strings.Contains(path, "/healthz") {
		// 	logmanager.Infof("%s %s %s", c.Request.Proto, c.Request.Method, path)
		// }
	}
}

// RequestLogger (HTTP Middleware)
// RequestLogger is a port of the Ginrus middleware from gin-gonic/contrib, but will
// include the request uuid as well.
func RequestLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Setting up dimensions for request logger
		start := time.Now()
		path := c.Request.URL.Path

		//Next sexecutes the pending handlers in the chain inside the calling handler.
		// Logic for response phase come over here
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		end = end.UTC()

		// Getting traceID from the contextkey set during the traceID middleware
		traceID, _ := c.Get(ContextKey)

		entry := logger.WithFields(logrus.Fields{
			"status":   c.Writer.Status(),
			"ip":       c.ClientIP(),
			"latency":  latency,
			ContextKey: traceID,
		})

		entry.Logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: LogTimestampFormat,
			FullTimestamp: true, ForceColors: true, DisableSorting: true, DisableLevelTruncation: true,
			SortingFunc: func(s []string) {
				sort.Strings(s)
			},
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			if !strings.Contains(path, "/healthz") {
				entry.Infof("%s %s %s %s", c.Request.Proto, http.StatusText(c.Writer.Status()), c.Request.Method, path)
			}
		}
	}
}

// Get the context aware logger for the inititiated gin context
// Handler get the logmanager context from the gin context
func GetContextAwareLogger(c *gin.Context) *logrus.Entry {
	logger, ok := c.Get(Contextawarelogger)
	if !ok {
		return logrus.StandardLogger().WithField(ContextKey, uuid.New().String())
	}
	Logger = logger.(*logrus.Entry)
	return Logger
}

func Get() *logrus.Entry {
	if Logger != nil {
		return Logger
	}

	logger := Initialize().WithFields(logrus.Fields{})
	return logger
}

// Get the traceID associated with current request
func GetTraceIDForRequest(c *gin.Context) string {
	traceID, _ := c.Get(ContextKey)
	return traceID.(string)
}

// Get the lowest logging level configuration
// If nothing specified or wrong value specified it will be default
func getLowestLoggingLevel() logrus.Level {

	lvl := utility.GetEnv("LOWEST_LOGGING_LEVEL", "debug")

	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}

	return ll
}
