package router

import (
	"github.com/airabinovich/memequotes_back/character"
	"github.com/airabinovich/memequotes_back/phrase"
	"github.com/gin-gonic/gin"
)

func mappings(router *gin.Engine) {
	router.POST("character", character.SaveCharacter)
	router.GET("characters", character.GetAllCharacters)
	router.GET("character/:character-id", character.GetCharacter)
	router.PATCH("character/:character-id", character.UpdateCharacter)
	router.DELETE("character/:character-id", character.DeleteCharacter)

	router.POST("character/:character-id/phrase", phrase.SaveNewPhrase)
	router.GET("character/:character-id/phrases", phrase.GetAllPhrasesForCharacter)
	router.GET("character/:character-id/phrase/:phrase-id", phrase.GetPhrase)
	router.DELETE("character/:character-id/phrase/:phrase-id", phrase.DeletePhraseForCharacter)
}
