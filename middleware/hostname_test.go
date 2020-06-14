package middleware

import (
	"context"
	"github.com/airabinovich/memequotes_back/utils"
	"net/http"
	"testing"

	ctx "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHostname(t *testing.T) {
	t.Log("Hostname should not be an empty string")

	gin.SetMode(gin.TestMode)
	hostname := ""
	router := gin.New()
	router.Use(Hostname)
	router.Use(func(c *gin.Context) {
		hostname = ctx.Hostname(ctx.RequestContext(c))
		c.Next()
	})

	utils.PerformRequest(router, http.MethodGet, "/", nil)

	assert.NotEmpty(t, hostname)
	assert.NotEqual(t, "undefined-hostname", hostname)
}

func TestHostnameWithoutRequestContext(t *testing.T) {
	t.Log("Hostname should be added to context")

	c := context.Background()
	c = HostnameWithoutRequestContext(c)

	hostname := ctx.Hostname(c)

	assert.NotEmpty(t, hostname)
	assert.NotEqual(t, "unknown", hostname)
}
