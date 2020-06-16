package phrase

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/database"
	"github.com/gin-gonic/gin"
	"time"
)

// GetAllForCharacter retrieves all phrases from a character
func GetAllForCharacter(c *gin.Context, characterId int64) ([]Phrase, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Phrase with characterId %d", characterId))

	phrases := make([]Phrase,0)
	db := database.DB.Where("character_id = ?", characterId).Find(&phrases)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return nil, false, db.Error
	}
	return phrases, !notFound, nil
}

// Save stores a new phrase for a character
func Save(c *gin.Context, phCmd PhraseCommand) (Phrase, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Creating Phrase for character %d", phCmd.CharacterId))

	now := time.Now()
	phrase := NewPhrase(0, phCmd.CharacterId, nil, phCmd.Content, now, now)
	if !database.DB.NewRecord(phrase) {
		return Phrase{}, errors.New("phrase already exists")
	}
	if err := database.DB.Create(&phrase).Error; err != nil {
		logger.Error("creating phrase", err)
		return Phrase{}, err
	}
	return phrase, nil
}
