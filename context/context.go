package context

import (
	"context"
	"github.com/airabinovich/memequotes_back/logger"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const (
	contextKey   = ctxKey("context_key")
	loggerKey    = ctxKey("logger_key")
	requestIDKey = ctxKey("request_id_key")
	hostnameKey  = ctxKey("hostname_key")
)

func (c ctxKey) String() string {
	return "back_context_key_" + string(c)
}

//WithRequestContext returns a context.Context from a gin.Context
func WithRequestContext(ctx context.Context, ginCtx *gin.Context) {
	ginCtx.Set(contextKey.String(), ctx)
}

//RequestContext returns a context.Context from a gin.Context
func RequestContext(ginCtx *gin.Context) context.Context {
	ctxValue, ok := ginCtx.Get(contextKey.String())
	if !ok {
		return context.Background()
	}

	return ctxValue.(context.Context)
}


//WithLogger stores a CreditsLogger in context
func WithLogger(ctx context.Context, logger *logger.SupportLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

//Logger returns the CreditsLogger from context
func Logger(ctx context.Context) *logger.SupportLogger {
	retrievedLogger, exists := ctx.Value(loggerKey).(*logger.SupportLogger)
	if !exists {
		retrievedLogger = logger.NewLogger(make(map[string]interface{}))
	}
	return retrievedLogger
}

// WithRequestID adds a RequestID to request context
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestID gets the request's requestID
func RequestID(ctx context.Context) string {
	uuid, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return "undefined-request_id"
	}
	return uuid
}

// WithHostname adds the hostname to request context
func WithHostname(ctx context.Context, hostname string) context.Context {
	return context.WithValue(ctx, hostnameKey, hostname)
}

// Hostname gets the hostname from request context
func Hostname(ctx context.Context) string {
	hostname, ok := ctx.Value(hostnameKey).(string)
	if !ok {
		return "undefined-hostname"
	}
	return hostname
}

// WithContext sets the application context
func WithContext(ctx context.Context, c context.Context) context.Context {
	return context.WithValue(ctx, contextKey, c)
}

// AppContext retrieves application context
func AppContext(ctx context.Context) context.Context {
	ctxValue, ok := ctx.Value(contextKey).(context.Context)
	if !ok {
		return context.Background()
	}
	return ctxValue
}
