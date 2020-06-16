package middleware

import (
	"github.com/airabinovich/memequotes_back/utils"
	"net/http"
	"testing"

	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestIDShouldBeAdded(t *testing.T) {
	t.Log("Request id should be added to context in every request")
	gin.SetMode(gin.TestMode)
	requestID := ""
	router := gin.New()
	router.Use(RequestID)
	router.Use(func(c *gin.Context) {
		ctx := commonContext.RequestContext(c)
		requestID = commonContext.RequestID(ctx)
		c.Next()
	})

	utils.PerformRequest(router, http.MethodGet, "/", nil)

	assert.NotEmpty(t, requestID)
	assert.NotEqual(t, "undefined-request_id", requestID)
}
