package character

import (
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/database"
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var characterRepository CharacterRepository

func Initialize() {
	characterRepository = NewDBCharacterRepository(database.DB)
}

func GetCharacter(c *gin.Context) {
	rest.ErrorWrapper(getCharacter, c)
}

func getCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric id", err)
		return rest.NewBadRequest(err.Error())
	}

	logger.Debug(fmt.Sprintf("Getting character with id %d", id))
	ch, found, err := characterRepository.Get(c, id)
	if err != nil {
		logger.Error("get character by id", err)
		return rest.NewInternalServerError(err.Error())
	}
	if !found {
		return rest.NewResourceNotFound(fmt.Sprintf("character %d not found", id))
	}

	c.JSON(http.StatusOK, CharacterResultFromCharacter(ch))
	return nil
}

// SaveCharacter saves a new character
func SaveCharacter(c *gin.Context) {
	rest.ErrorWrapper(saveCharacter, c)
}

func saveCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	var chCmd CharacterCommand
	if err := c.ShouldBindJSON(&chCmd); err != nil {
		logger.Error("creating character bad body format", err)
		return rest.NewBadRequest(err.Error())
	}

	ch, err := characterRepository.Save(c, chCmd)
	if err != nil {
		logger.Error("error creating character", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, CharacterResultFromCharacter(ch))
	return nil
}

// UpdateCharacter updates an existing character
func UpdateCharacter(c *gin.Context) {
	rest.ErrorWrapper(updateCharacter, c)
}

func updateCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric id", err)
		return rest.NewBadRequest(err.Error())
	}

	var chCmd CharacterCommand
	if err := c.ShouldBindJSON(&chCmd); err != nil {
		logger.Error("updating character bad body format", err)
		return rest.NewBadRequest(err.Error())
	}

	logger.Debug(fmt.Sprintf("Updating character with id %d", id))
	ch, found, err := characterRepository.Update(c, id, chCmd)
	if err != nil {
		logger.Error("update character by id", err)
		return rest.NewInternalServerError(err.Error())
	}
	if !found {
		return rest.NewResourceNotFound(fmt.Sprintf("character %d not found", id))
	}

	c.JSON(http.StatusOK, CharacterResultFromCharacter(ch))
	return nil
}
