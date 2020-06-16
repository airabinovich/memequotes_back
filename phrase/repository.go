package phrase

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	customErrors "github.com/airabinovich/memequotes_back/errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type PhraseRepository interface {
	// Get a phrase for a character
	Get(c *gin.Context, characterId int64, id int64) (Phrase, bool, error)

	// GetAllForCharacter retrieves all phrases from a character
	GetAllForCharacter(c *gin.Context, characterId int64) ([]Phrase, bool, error)

	// Save stores a new phrase for a character
	Save(c *gin.Context, phCmd PhraseCommand) (Phrase, error)

	// Delete a phrase for a character
	Delete(c *gin.Context, characterId int64, id int64) error
}

type DBPhraseRepository struct {
	db *gorm.DB
}

func NewDBPhraseRepository(db *gorm.DB) DBPhraseRepository {
	return DBPhraseRepository{
		db: db,
	}
}

func (repo DBPhraseRepository) Get(c *gin.Context, characterId int64, id int64) (Phrase, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Phrase with characterId %d and id %d", characterId, id))

	phrase := Phrase{}
	db := repo.db.Where("id = ?", id).First(&phrase)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return Phrase{}, false, db.Error
	}

	if notFound {
		return Phrase{}, false, nil
	}

	if phrase.CharacterId != characterId {
		return Phrase{}, false, customErrors.NewUnauthorizedError("phrase doesn't belong to character")
	}

	return phrase, !notFound, nil
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

func (repo DBPhraseRepository) Delete(c *gin.Context, characterId int64, id int64) error {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Phrase with characterId %d and id %d", characterId, id))

	phrase, found, err := repo.Get(c, characterId, id)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}

	db := repo.db.Delete(&phrase)
	if db.Error != nil {
		return db.Error
	}
	return nil
}