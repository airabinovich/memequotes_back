package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(r http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	return router
}

