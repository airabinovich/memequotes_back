package middleware

import (
	"context"

	commonContext "github.com/airabinovich/memequotes_back/context"
)

// NoRequestContext sets up a new application context based on a context
func NoRequestContext(c context.Context) context.Context {
	appCtx := commonContext.AppContext(c)
	appCtx = HostnameWithoutRequestContext(appCtx)
	appCtx = RequestIDWithNoRequestContext(appCtx)
	appCtx = LoggerWithoutRequestContext(appCtx)
	return commonContext.WithContext(c, appCtx)
}

// RefreshRequestIDContext refresh requestID in application context
func RefreshRequestIDContext(c context.Context) context.Context {
	appCtx := commonContext.AppContext(c)
	appCtx = RequestIDWithNoRequestContext(appCtx)
	appCtx = LoggerWithoutRequestContext(appCtx)
	appCtx = commonContext.WithContext(c, appCtx)
	return appCtx
}

