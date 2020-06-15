package character

import (
	"fmt"
	"github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCharacter(c *gin.Context) {
	rest.ErrorWrapper(getCharacter, c)
}

func getCharacter(c *gin.Context) *rest.APIError {
	ctx := context.RequestContext(c)
	logger := context.Logger(ctx)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric id", err)
		return rest.NewBadRequest(err.Error())
	}

	logger.Debug(fmt.Sprintf("Getting character with id %d", id))
	ch, found, err := Get(c, id)
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

func SaveCharacter(c *gin.Context) {
	rest.ErrorWrapper(saveCharacter, c)
}

func saveCharacter(c *gin.Context) *rest.APIError {
	ctx := context.RequestContext(c)
	logger := context.Logger(ctx)

	var chCmd CharacterCommand
	if err := c.ShouldBindJSON(&chCmd); err != nil {
		logger.Error("creating character bad body format", err)
		return rest.NewBadRequest(err.Error())
	}

	ch, err := Save(c, chCmd)
	if err != nil {
		logger.Error("error creating character", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, CharacterResultFromCharacter(ch))
	return nil
}
