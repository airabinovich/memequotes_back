package router

import (
	"github.com/airabinovich/memequotes_back/character"
	"github.com/airabinovich/memequotes_back/handlers"
	"github.com/airabinovich/memequotes_back/phrase"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine) {

	router.GET("miami", handlers.ComandateHandler)

	router.POST("character", character.SaveCharacter)
	router.GET("character/:id", character.GetCharacter)

	router.POST("character/:id/phrase", phrase.SaveNewPhrase)
	router.GET("character/:id/phrases", phrase.GetAllPhrasesForCharacter)
}
