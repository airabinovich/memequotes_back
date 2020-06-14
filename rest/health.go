package rest

import (
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health controller for checking service health
func Health(c *gin.Context) {
	ErrorWrapper(health, c)
}

func health(c *gin.Context) *APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Trace("health-check")

	c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{
		"I'm Ok",
	})
	return nil
}

