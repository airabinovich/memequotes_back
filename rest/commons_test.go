package rest

import (
	"github.com/airabinovich/memequotes_back/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	t.Log("Method not allowed handler should be called and return 405")

	router := utils.TestRouter()
	router.NoMethod(MethodNotAllowedHandler)
	router.NoRoute(NoRouteHandler)
	router.GET("/health", Health)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestNoRouteHandler(t *testing.T) {
	t.Log("No route handler should be called and return 404 when there's no method to resolve the given request")

	router := utils.TestRouter()
	router.NoMethod(MethodNotAllowedHandler)
	router.NoRoute(NoRouteHandler)
	router.GET("/health", Health)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}