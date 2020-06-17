package phrase

import (
	"fmt"
	commonContext "github.com/airabinovich/memequotes_back/context"
	customErrors "github.com/airabinovich/memequotes_back/errors"
	"github.com/airabinovich/memequotes_back/model"
	"github.com/airabinovich/memequotes_back/repository"
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var phraseRepository repository.PhraseRepository

func Initialize(phRepo repository.PhraseRepository) {
	phraseRepository = phRepo
}

func GetPhrase(c *gin.Context) {
	rest.ErrorWrapper(getPhrase, c)
}

func getPhrase(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	phraseId, err := strconv.ParseInt(c.Param("phrase-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	phrase, found, err := phraseRepository.Get(c, characterId, phraseId)
	if err != nil {
		switch err.(type) {
		case customErrors.UnauthorizedError:
			return rest.NewUnauthorized(err.Error())
		default:
			return rest.NewInternalServerError(err.Error())
		}
	}

	if !found {
		return rest.NewResourceNotFound("phrase not found")
	}

	c.JSON(http.StatusOK, model.PhraseResultFromPhrase(phrase))
	return nil
}

// GetAllPhrasesForCharacter returns all phrases for a character wrapped in a json object
func GetAllPhrasesForCharacter(c *gin.Context) {
	rest.ErrorWrapper(getAllPhrasesForCharacter, c)
}

func getAllPhrasesForCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	logger.Debug(fmt.Sprintf("Getting phrases for character id %d", characterId))
	phrases, found, err := phraseRepository.GetAllForCharacter(c, characterId)
	if err != nil {
		logger.Error("get character by id", err)
		return rest.NewInternalServerError(err.Error())
	}
	if !found {
		return rest.NewResourceNotFound(fmt.Sprintf("phrases for character %d not found", characterId))
	}
	phraseResults := make([]model.PhraseResult, len(phrases))
	for i, phrase := range phrases {
		phraseResults[i] = model.PhraseResultFromPhrase(phrase)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"results": phraseResults,
	})
	return nil
}

// SaveNewPhrase saves a new phrase for the specified character
func SaveNewPhrase(c *gin.Context) {
	rest.ErrorWrapper(saveNewPhrase, c)
}

func saveNewPhrase(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	var phCmd model.PhraseCommand
	if err := c.ShouldBindJSON(&phCmd); err != nil {
		logger.Error("creating phrase bad body format", err)
		return rest.NewBadRequest(err.Error())
	}
	phCmd.CharacterId = characterId

	phrase, err := phraseRepository.Save(c, phCmd)
	if err != nil {
		logger.Error("error creating phrase", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, model.PhraseResultFromPhrase(phrase))
	return nil
}

func DeletePhraseForCharacter(c *gin.Context) {
	rest.ErrorWrapper(deletePhraseForCharacter, c)
}

func deletePhraseForCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	id, err := strconv.ParseInt(c.Param("phrase-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric id", err)
		return rest.NewBadRequest(err.Error())
	}

	err = phraseRepository.Delete(c, characterId, id)
	if err != nil {
		switch err.(type) {
		case customErrors.UnauthorizedError:
			return rest.NewUnauthorized(err.Error())
		default:
			return rest.NewInternalServerError(err.Error())
		}
	}

	c.Status(http.StatusGone)
	return nil
}