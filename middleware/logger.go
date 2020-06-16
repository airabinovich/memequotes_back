package middleware

import (
	"context"

	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/logger"
	"github.com/gin-gonic/gin"
)

//Logger sets up a new logger with request information
func Logger(c *gin.Context) {
	requestCtx := commonContext.RequestContext(c)

	fields := make(map[string]interface{})
	fields["x-request-id"] = commonContext.RequestID(requestCtx)
	fields["hostname"] = commonContext.Hostname(requestCtx)

	l := logger.NewLogger(fields)

	requestCtx = commonContext.WithLogger(requestCtx, l)
	commonContext.WithRequestContext(requestCtx, c)
	c.Next()
}

// LoggerWithoutRequestContext sets up a new logger with application context
func LoggerWithoutRequestContext(c context.Context) context.Context {
	fields := make(map[string]interface{})
	fields["x-request-id"] = commonContext.RequestID(c)
	fields["hostname"] = commonContext.Hostname(c)

	l := logger.NewLogger(fields)

	return commonContext.WithLogger(c, l)
}
