package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthShouldReturn200Ok(t *testing.T) {
	t.Log("Health should return 200 OK when service is healthy")
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/memequotes/health", nil)
	r := testRouter()
	r.GET("memequotes/health", Health)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// TODO: this could be improved, here we are creating a new router. We cannot use router.CreateRouter() because it generates cyclic references
func testRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	return router
}
