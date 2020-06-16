package phrase

import (
	"fmt"
	"github.com/airabinovich/memequotes_back/context"
	"github.com/airabinovich/memequotes_back/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllPhrasesForCharacter returns all phrases for a character wrapped in a json object
func GetAllPhrasesForCharacter(c *gin.Context) {
	rest.ErrorWrapper(getAllPhrasesForCharacter, c)
}

func getAllPhrasesForCharacter(c *gin.Context) *rest.APIError {
	ctx := context.RequestContext(c)
	logger := context.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	logger.Debug(fmt.Sprintf("Getting phrases for character id %d", characterId))
	phrases, found, err := GetAllForCharacter(c, characterId)
	if err != nil {
		logger.Error("get character by id", err)
		return rest.NewInternalServerError(err.Error())
	}
	if !found {
		return rest.NewResourceNotFound(fmt.Sprintf("phrases for character %d not found", characterId))
	}
	phraseResults := make([]PhraseResult, len(phrases))
	for i, phrase := range phrases {
		phraseResults[i] = PhraseResultFromPhrase(phrase)
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
	ctx := context.RequestContext(c)
	logger := context.Logger(ctx)

	var phCmd PhraseCommand
	if err := c.ShouldBindJSON(&phCmd); err != nil {
		logger.Error("creating phrase bad body format", err)
		return rest.NewBadRequest(err.Error())
	}

	phrase, err := Save(c, phCmd)
	if err != nil {
		logger.Error("error creating phrase", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, PhraseResultFromPhrase(phrase))
	return nil
}