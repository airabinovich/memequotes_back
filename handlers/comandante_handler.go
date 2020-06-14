package handlers

import (
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

var msgs = [...]string{
	"en mi rolex falta un minuto todavía",
	"yo no manejo el rating, manejo un Rolls Royce",
	"sacá la mano de ahí carajo!",
	"te va queda letrificada loca!",
	"el tren de Ricardo Fort pasa una sola vez en la vida",
	"maiameeeeee",
	"basta chicos",
}

func ComandateHandler(ctx *gin.Context) {
	rest.ErrorWrapper(comandanteHandler, ctx)
}

func comandanteHandler(c *gin.Context) *rest.APIError {
	logger := commonContext.Logger(c)
	logger.Info("Received request in Miami")

	nextIndex := rand.Int() % len(msgs)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": msgs[nextIndex],
	})

	return nil
}
