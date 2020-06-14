package middleware

import (
	"context"

	ctx "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// RequestID adds request-id to the request context.
func RequestID(c *gin.Context) {
	requestCtx := ctx.WithRequestID(ctx.RequestContext(c), uuid.NewV4().String())
	ctx.WithRequestContext(requestCtx, c)
	c.Next()
}

// RequestIDWithNoRequestContext sets up request-id to application context
func RequestIDWithNoRequestContext(c context.Context) context.Context {
	return ctx.WithRequestID(c, uuid.NewV4().String())
}

