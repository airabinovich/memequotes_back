package repository

import (
	"github.com/airabinovich/memequotes_back/model"
	"github.com/gin-gonic/gin"
)

type PhraseRepository interface {
	// Get a phrase for a character
	Get(c *gin.Context, characterId int64, id int64) (model.Phrase, bool, error)

	// GetAllForCharacter retrieves all phrases from a character
	GetAllForCharacter(c *gin.Context, characterId int64) ([]model.Phrase, bool, error)

	// Save stores a new phrase for a character
	Save(c *gin.Context, phCmd model.PhraseCommand) (model.Phrase, error)

	// Delete a phrase for a character
	Delete(c *gin.Context, characterId int64, id int64) error
}
