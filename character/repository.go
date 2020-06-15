package character

import (
	"errors"
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/database"
	"github.com/gin-gonic/gin"
	"time"
)

//Get a Character by id. Returns the character, whether it's found and an error
func Get(c *gin.Context, id int64) (Character, bool, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Getting Character with id %d", id))

	ch := Character{}
	db := database.DB.Where("id = ?", id).Find(&ch)
	notFound := db.RecordNotFound()
	if db.Error != nil && !notFound {
		return Character{}, false, db.Error
	}
	return ch, !notFound, nil
}

func Save(c *gin.Context, chCmd CharacterCommand) (Character, error) {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)
	logger.Debug(fmt.Sprintf("Creating Character with name %s", chCmd.Name))

	now := time.Now()
	ch := NewCharacter(0, chCmd.Name, now, now)
	fmt.Printf("SAVING CHARACTER %v", ch)
	if !database.DB.NewRecord(ch) {
		return Character{}, errors.New("characters already exists")
	}
	if err := database.DB.Create(&ch).Error; err != nil {
		logger.Error("creating character", err)
		return Character{}, err
	}
	return ch, nil
}