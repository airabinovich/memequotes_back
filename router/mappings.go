package router

import (
	"github.com/airabinovich/memequotes_back/handlers"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine) {

	router.GET("miami", handlers.ComandateHandler)
}
