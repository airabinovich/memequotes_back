package repository

import (
	"github.com/airabinovich/memequotes_back/model"
	"github.com/gin-gonic/gin"
)

type CharacterRepository interface {
	// Get a Character by id. Returns the character, whether it's found and an error
	Get(c *gin.Context, id int64) (model.Character, bool, error)

	// GetAll retrieves all character in the repository
	GetAll(c *gin.Context) ([]model.Character, error)

	// Save stores a new character
	Save(c *gin.Context, chCmd model.CharacterCommand) (model.Character, error)

	// Update a character. Returns the updated character, whether it's found and an error
	Update(c *gin.Context, id int64, chCmd model.CharacterCommand) (model.Character, bool, error)

	// Delete a character. If the character has phrases this will fail. Remove all phrases before
	Delete(c *gin.Context, id int64) error
}
