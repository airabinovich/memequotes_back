package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// NoRouteHandler handles requests for non registered routes
func NoRouteHandler(c *gin.Context) {
	ErrorWrapper(func(c *gin.Context) *APIError {
		return NewResourceNotFound(fmt.Sprintf("Resource not found for %s.", c.Request.URL.Path))
	}, c)
}

// MethodNotAllowedHandler handles requests for registered routes with invalid http methods on their requests
func MethodNotAllowedHandler(c *gin.Context) {
	ErrorWrapper(func(c *gin.Context) *APIError {
		return NewMethodNotAllowed("Method not allowed - %s - %s ", c.Request.Method, c.Request.URL.Path)
	}, c)
}

// WrapperFunc is the func type for the custom handlers.
type WrapperFunc func(c *gin.Context) *APIError

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func ErrorWrapper(handlerFunc WrapperFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

