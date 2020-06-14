package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRouter(t *testing.T) {
	t.Log("CreateRouter should return a new router")

	r := CreateRouter()

	assert.NotNil(t, r)
	assert.True(t, r.HandleMethodNotAllowed)
}
