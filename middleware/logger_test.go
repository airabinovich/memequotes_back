package middleware

import (
	"context"
	"github.com/airabinovich/memequotes_back/utils"
	"testing"

	commonContext "github.com/airabinovich/memequotes_back/context"
	log "github.com/airabinovich/memequotes_back/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoggerWithRequestID(t *testing.T) {
	t.Log("When no request-id is present in context, a new logger should be created with a new value for request_id and no value for hostname")

	var logger *log.SupportLogger
	gin.SetMode(gin.TestMode)
	router := gin.New()
	requestID := ""

	router.Use(RequestID)
	router.Use(func(c *gin.Context) {
		requestID = commonContext.RequestID(commonContext.RequestContext(c))
		c.Next()
	})
	router.Use(Logger)
	router.Use(func(c *gin.Context) {
		logger = commonContext.Logger(commonContext.RequestContext(c))
		c.Next()
	})

	utils.PerformRequest(router, "GET", "/", map[string]string{})

	assert.NotNil(t, logger)
	assert.Equal(t, log.NewLogger(map[string]interface{}{"x-request-id": requestID, "hostname": "undefined-hostname"}), logger)
}

func TestLoggerWithRequestIDAndHostname(t *testing.T) {
	t.Log("When no request-id is present in context, a new logger should be created with a new value for request_id and value for hostname")

	var logger *log.SupportLogger
	gin.SetMode(gin.TestMode)
	router := gin.New()
	requestID := ""
	hostname := ""

	router.Use(RequestID)
	router.Use(Hostname)
	router.Use(func(c *gin.Context) {
		reqCtx := commonContext.RequestContext(c)
		requestID = commonContext.RequestID(reqCtx)
		hostname = commonContext.Hostname(reqCtx)
		c.Next()
	})
	router.Use(Logger)
	router.Use(func(c *gin.Context) {
		logger = commonContext.Logger(commonContext.RequestContext(c))
		c.Next()
	})

	utils.PerformRequest(router, "GET", "/", map[string]string{})

	assert.NotNil(t, logger)
	assert.Equal(t, log.NewLogger(map[string]interface{}{"x-request-id": requestID, "hostname": hostname}), logger)
}

func TestLoggerWithHostnameAndWithoutRequestID(t *testing.T) {
	t.Log("When no request-id is present in context, a new logger should be created with a value for hostname and no value for request-id")

	var logger *log.SupportLogger
	gin.SetMode(gin.TestMode)
	router := gin.New()
	hostname := ""

	router.Use(Hostname)
	router.Use(func(c *gin.Context) {
		hostname = commonContext.Hostname(commonContext.RequestContext(c))
		c.Next()
	})
	router.Use(Logger)
	router.Use(func(c *gin.Context) {
		logger = commonContext.Logger(commonContext.RequestContext(c))
		c.Next()
	})

	utils.PerformRequest(router, "GET", "/", map[string]string{})

	assert.NotNil(t, logger)
	assert.Equal(t, log.NewLogger(map[string]interface{}{"x-request-id": "undefined-request_id", "hostname": hostname}), logger)
}

func TestLoggerWithoutRequestContext(t *testing.T) {
	t.Log("Logger should be created with application context. This includes hostname and x-request-id")

	var logger *log.SupportLogger
	c := context.Background()
	c = commonContext.AppContext(NoRequestContext(c))

	hostname := commonContext.Hostname(c)
	requestID := commonContext.RequestID(c)

	logger = commonContext.Logger(c)

	assert.NotNil(t, logger)
	assert.Equal(t, requestID, logger.Data["x-request-id"])
	assert.Equal(t, hostname, logger.Data["hostname"])
}

func TestLoggerWithoutRequestContextAndRefreshingRequestID(t *testing.T) {
	t.Log("Logger should be created with application context. This includes hostname and x-request-id. Then refresh request-id")

	var logger *log.SupportLogger
	c := context.Background()
	c = NoRequestContext(c)
	appCtx := commonContext.AppContext(c)

	hostname := commonContext.Hostname(appCtx)
	initialRequestID := commonContext.RequestID(appCtx)

	appCtx = commonContext.AppContext(RefreshRequestIDContext(c))
	refreshedRequestID := commonContext.RequestID(appCtx)

	logger = commonContext.Logger(appCtx)

	assert.NotNil(t, logger)
	assert.NotEqual(t, initialRequestID, logger.Data["x-request-id"])
	assert.Equal(t, refreshedRequestID, logger.Data["x-request-id"])
	assert.Equal(t, hostname, logger.Data["hostname"])
}

