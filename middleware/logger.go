package middleware

import (
	"context"

	ctx "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/logger"
	"github.com/gin-gonic/gin"
)

//Logger sets up a new logger with request information
func Logger(c *gin.Context) {
	requestCtx := ctx.RequestContext(c)

	fields := make(map[string]interface{})
	fields["x-request-id"] = ctx.RequestID(requestCtx)
	fields["hostname"] = ctx.Hostname(requestCtx)

	l := logger.NewLogger(fields)

	requestCtx = ctx.WithLogger(requestCtx, l)
	ctx.WithRequestContext(requestCtx, c)
	c.Next()
}

// LoggerWithoutRequestContext sets up a new logger with application context
func LoggerWithoutRequestContext(c context.Context) context.Context {
	fields := make(map[string]interface{})
	fields["x-request-id"] = ctx.RequestID(c)
	fields["hostname"] = ctx.Hostname(c)

	l := logger.NewLogger(fields)

	return ctx.WithLogger(c, l)
}
