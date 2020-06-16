package character

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type CharacterRepository interface {
	// Get a Character by id. Returns the character, whether it's found and an error
	Get(c *gin.Context, id int64) (Character, bool, error)

	// Save stores a new character
	Save(c *gin.Context, chCmd CharacterCommand) (Character, error)

	// Update a character. Returns the updated character, whether it's found and an error
	Update(c *gin.Context, id int64, chCmd CharacterCommand) (Character, bool, error)
}

type DBCharacterRepository struct {
	db *gorm.DB
}

func NewDBCharacterRepository(db *gorm.DB) DBCharacterRepository {
	return DBCharacterRepository{
		db: db,
	}
}

func (repo DBCharacterRepository) Get(c *gin.Context, id int64) (Character, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Character with id %d", id))

	ch := Character{}
	db := repo.db.Where("id = ?", id).Find(&ch)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return Character{}, false, db.Error
	}
	return ch, !notFound, nil
}

func (repo DBCharacterRepository) Save(c *gin.Context, chCmd CharacterCommand) (Character, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Creating Character with name %s", chCmd.Name))

	now := time.Now()
	ch := NewCharacter(0, chCmd.Name, now, now)
	if !repo.db.NewRecord(ch) {
		return Character{}, errors.New("characters already exists")
	}
	if err := repo.db.Create(&ch).Error; err != nil {
		logger.Error("creating character", err)
		return Character{}, err
	}
	return ch, nil
}

func (repo DBCharacterRepository) Update(c *gin.Context, id int64, chCmd CharacterCommand) (Character, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Updating Character with id %d", id))

	ch, found, err := repo.Get(c, id)
	if err != nil {
		logger.Error("error retrieving character", err)
		return Character{}, false, err
	}
	if !found {
		return Character{}, false, nil
	}

	ch.Name = chCmd.Name
	ch.LastUpdated = time.Now()
	if err := repo.db.Save(&ch).Error; err != nil {
		logger.Error("updating character", err)
		return Character{}, true, err
	}

	return ch, true, nil
}