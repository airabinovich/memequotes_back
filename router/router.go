package router

import (
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
)

// Route creates a new router
func Route() *gin.Engine {
	router := rest.CreateRouter()
	mappings(router)
	return router
}
