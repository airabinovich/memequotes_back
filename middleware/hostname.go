package middleware

import (
	"context"
	"os"

	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
)

// Hostname gets the hostname from os and adds the hostname to request context
func Hostname(c *gin.Context) {
	withCtx := commonContext.WithHostname(commonContext.RequestContext(c), hostname())
	commonContext.WithRequestContext(withCtx, c)
	c.Next()
}

// HostnameWithoutRequestContext gets the hostname os and add it to app context
func HostnameWithoutRequestContext(c context.Context) context.Context {
	return commonContext.WithHostname(c, hostname())
}

func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return hostname
}
