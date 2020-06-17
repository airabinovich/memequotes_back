package character

import (
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/model"
	"github.com/airabinovich/memequotes_back/repository"
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var characterRepository repository.CharacterRepository
var phraseRepository repository.PhraseRepository

func Initialize(chRepo repository.CharacterRepository, phRepo repository.PhraseRepository) {
	characterRepository = chRepo
	phraseRepository = phRepo
}

func GetCharacter(c *gin.Context) {
	rest.ErrorWrapper(getCharacter, c)
}

func getCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	id, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
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

	c.JSON(http.StatusOK, model.CharacterResultFromCharacter(ch))
	return nil
}

// SaveCharacter saves a new character
func SaveCharacter(c *gin.Context) {
	rest.ErrorWrapper(saveCharacter, c)
}

func saveCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	var chCmd model.CharacterCommand
	if err := c.ShouldBindJSON(&chCmd); err != nil {
		logger.Error("creating character bad body format", err)
		return rest.NewBadRequest(err.Error())
	}

	ch, err := characterRepository.Save(c, chCmd)
	if err != nil {
		logger.Error("error creating character", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, model.CharacterResultFromCharacter(ch))
	return nil
}

func GetAllCharacters(c *gin.Context) {
	rest.ErrorWrapper(getAllCharacters, c)
}

func getAllCharacters(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	logger.Debug("Getting all characters")
	chs, err := characterRepository.GetAll(c)
	if err != nil {
		logger.Error("get all character", err)
		return rest.NewInternalServerError(err.Error())
	}

	chResults := make([]model.CharacterResult, len(chs))
	for i, ch := range chs {
		chResults[i] = model.CharacterResultFromCharacter(ch)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"results": chResults,
	})
	return nil
}

// UpdateCharacter updates an existing character
func UpdateCharacter(c *gin.Context) {
	rest.ErrorWrapper(updateCharacter, c)
}

func updateCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	id, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric id", err)
		return rest.NewBadRequest(err.Error())
	}

	var chCmd model.CharacterCommand
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

	c.JSON(http.StatusOK, model.CharacterResultFromCharacter(ch))
	return nil
}

func DeleteCharacter(c *gin.Context) {
	rest.ErrorWrapper(deleteCharacter, c)
}

func deleteCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric id", err)
		return rest.NewBadRequest(err.Error())
	}

	phrases, found, err := phraseRepository.GetAllForCharacter(c, characterId)
	if err != nil {
		logger.Error("cannot get phrases for character", err)
		return rest.NewInternalServerError(err.Error())
	}
	if found && len(phrases) > 0 {
		for _, ph := range phrases {
			if err := phraseRepository.Delete(c, characterId, ph.ID); err != nil {
				logger.Error("cannot delete phrase from character", err)
				return rest.NewInternalServerError(err.Error())
			}
		}
	}

	err = characterRepository.Delete(c, characterId)
	if err != nil {
		logger.Error("error deleting character", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.Status(http.StatusGone)
	return nil
}
