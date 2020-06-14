package middleware

import (
	"context"

	ctx "github.com/airabinovich/memequotes_back/context"
)

// NoRequestContext sets up a new application context based on a context
func NoRequestContext(c context.Context) context.Context {
	appCtx := ctx.AppContext(c)
	appCtx = HostnameWithoutRequestContext(appCtx)
	appCtx = RequestIDWithNoRequestContext(appCtx)
	appCtx = LoggerWithoutRequestContext(appCtx)
	return ctx.WithContext(c, appCtx)
}

// RefreshRequestIDContext refresh requestID in application context
func RefreshRequestIDContext(c context.Context) context.Context {
	appCtx := ctx.AppContext(c)
	appCtx = RequestIDWithNoRequestContext(appCtx)
	appCtx = LoggerWithoutRequestContext(appCtx)
	appCtx = ctx.WithContext(c, appCtx)
	return appCtx
}

