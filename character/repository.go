package character

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type DBCharacterRepository struct {
	db *gorm.DB
}

func NewDBCharacterRepository(db *gorm.DB) DBCharacterRepository {
	return DBCharacterRepository{
		db: db,
	}
}

func (repo DBCharacterRepository) Get(c *gin.Context, id int64) (model.Character, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Character with id %d", id))

	ch := model.Character{}
	db := repo.db.Where("id = ?", id).Find(&ch)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return model.Character{}, false, db.Error
	}
	return ch, !notFound, nil
}

func (repo DBCharacterRepository) GetAll(c *gin.Context) ([]model.Character, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug("Getting all Characters")

	chs := make([]model.Character, 0)
	db := repo.db.Find(&chs)
	if db.Error != nil {
		return []model.Character{}, db.Error
	}

	return chs, nil
}

func (repo DBCharacterRepository) Save(c *gin.Context, chCmd model.CharacterCommand) (model.Character, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Creating Character with name %s", chCmd.Name))

	now := time.Now()
	ch := model.NewCharacter(0, chCmd.Name, now, now)
	if !repo.db.NewRecord(ch) {
		return model.Character{}, errors.New("characters already exists")
	}
	if err := repo.db.Create(&ch).Error; err != nil {
		logger.Error("creating character", err)
		return model.Character{}, err
	}
	return ch, nil
}

func (repo DBCharacterRepository) Update(c *gin.Context, id int64, chCmd model.CharacterCommand) (model.Character, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Updating Character with id %d", id))

	ch, found, err := repo.Get(c, id)
	if err != nil {
		logger.Error("error retrieving character", err)
		return model.Character{}, false, err
	}
	if !found {
		return model.Character{}, false, nil
	}

	ch.Name = chCmd.Name
	ch.LastUpdated = time.Now()
	if err := repo.db.Save(&ch).Error; err != nil {
		logger.Error("updating character", err)
		return model.Character{}, true, err
	}

	return ch, true, nil
}

func (repo DBCharacterRepository) Delete(c *gin.Context, id int64) error {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Deleting Character with id %d", id))

	phrase, found, err := repo.Get(c, id)
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