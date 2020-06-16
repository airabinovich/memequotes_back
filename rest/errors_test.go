package rest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBadRequest(t *testing.T) {
	t.Log("New bad request should return a new bad request error")

	err := NewBadRequest("test message")

	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "test message", err.Message)
	assert.Equal(t, "bad_request", err.Err)
}

func TestNewInternalServerError(t *testing.T) {
	t.Log("NewInternalServerError should return a new internal server error")

	err := NewInternalServerError("some error")

	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "some error", err.Message)
	assert.Equal(t, "internal_error", err.Err)
}

func TestNewUnauthorized(t *testing.T) {
	t.Log("NewUnauthorized should return a new unauthorized")

	err := NewUnauthorized("some error")

	assert.Equal(t, http.StatusUnauthorized, err.Status)
	assert.Equal(t, "some error", err.Message)
	assert.Equal(t, "unauthorized", err.Err)
}