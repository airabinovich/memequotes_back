package rest

import (
	"github.com/airabinovich/memequotes_back/config"
	"github.com/airabinovich/memequotes_back/middleware"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	c := config.Conf
	env := c.GetString("environment", "dev")
	if env == "prod" || env == "beta" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.HandleMethodNotAllowed = true

	router.Use(middleware.Hostname)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.NoMethod(MethodNotAllowedHandler)
	router.NoRoute(NoRouteHandler)

	return router
}
