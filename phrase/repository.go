package phrase

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	customErrors "github.com/airabinovich/memequotes_back/errors"
	"github.com/airabinovich/memequotes_back/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type DBPhraseRepository struct {
	db *gorm.DB
}

func NewDBPhraseRepository(db *gorm.DB) DBPhraseRepository {
	return DBPhraseRepository{
		db: db,
	}
}

func (repo DBPhraseRepository) Get(c *gin.Context, characterId int64, id int64) (model.Phrase, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Phrase with characterId %d and id %d", characterId, id))

	phrase := model.Phrase{}
	db := repo.db.Where("id = ?", id).First(&phrase)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return model.Phrase{}, false, db.Error
	}

	if notFound {
		return model.Phrase{}, false, nil
	}

	if phrase.CharacterId != characterId {
		return model.Phrase{}, false, customErrors.NewUnauthorizedError("phrase doesn't belong to character")
	}

	return phrase, !notFound, nil
}

func (repo DBPhraseRepository) GetAllForCharacter(c *gin.Context, characterId int64) ([]model.Phrase, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Phrase with characterId %d", characterId))

	phrases := make([]model.Phrase, 0)
	db := repo.db.Where("character_id = ?", characterId).Find(&phrases)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return nil, false, db.Error
	}
	return phrases, !notFound, nil
}

func (repo DBPhraseRepository) Save(c *gin.Context, phCmd model.PhraseCommand) (model.Phrase, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Creating Phrase for character %d", phCmd.CharacterId))

	now := time.Now()
	phrase := model.NewPhrase(0, phCmd.CharacterId, nil, phCmd.Content, now, now)
	if !repo.db.NewRecord(phrase) {
		return model.Phrase{}, errors.New("phrase already exists")
	}
	if err := repo.db.Create(&phrase).Error; err != nil {
		logger.Error("creating phrase", err)
		return model.Phrase{}, err
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