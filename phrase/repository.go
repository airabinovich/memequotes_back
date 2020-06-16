package phrase

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type PhraseRepository interface {
	// GetAllForCharacter retrieves all phrases from a character
	GetAllForCharacter(c *gin.Context, characterId int64) ([]Phrase, bool, error)

	// Save stores a new phrase for a character
	Save(c *gin.Context, phCmd PhraseCommand) (Phrase, error)
}

type DBPhraseRepository struct {
	db *gorm.DB
}

func NewDBPhraseRepository(db *gorm.DB) DBPhraseRepository {
	return DBPhraseRepository{
		db: db,
	}
}

func (repo DBPhraseRepository) GetAllForCharacter(c *gin.Context, characterId int64) ([]Phrase, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Phrase with characterId %d", characterId))

	phrases := make([]Phrase, 0)
	db := repo.db.Where("character_id = ?", characterId).Find(&phrases)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return nil, false, db.Error
	}
	return phrases, !notFound, nil
}

func (repo DBPhraseRepository) Save(c *gin.Context, phCmd PhraseCommand) (Phrase, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Creating Phrase for character %d", phCmd.CharacterId))

	now := time.Now()
	phrase := NewPhrase(0, phCmd.CharacterId, nil, phCmd.Content, now, now)
	if !repo.db.NewRecord(phrase) {
		return Phrase{}, errors.New("phrase already exists")
	}
	if err := repo.db.Create(&phrase).Error; err != nil {
		logger.Error("creating phrase", err)
		return Phrase{}, err
	}
	return phrase, nil
}
