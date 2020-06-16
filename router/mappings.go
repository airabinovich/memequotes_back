package router

import (
	"github.com/airabinovich/memequotes_back/character"
	"github.com/airabinovich/memequotes_back/phrase"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine) {
	router.POST("character", character.SaveCharacter)
	router.GET("character/:id", character.GetCharacter)
	router.PATCH("character/:id", character.UpdateCharacter)

	router.POST("character/:id/phrase", phrase.SaveNewPhrase)
	router.GET("character/:id/phrases", phrase.GetAllPhrasesForCharacter)
	router.GET("character/:id/phrase/:phrase-id", phrase.GetPhrase)
	router.DELETE("character/:character-id/phrase/:id", phrase.DeletePhraseForCharacter)
}
