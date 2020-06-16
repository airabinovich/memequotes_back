package phrase

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/airabinovich/memequotes_back/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetPhrasesWithNonNumericIdShouldFail(t *testing.T) {
	t.Log("Calling with a non-numeric ID should return an error")

	w := httptest.NewRecorder()


	req := httptest.NewRequest(http.MethodGet, "/character/john/phrases", nil)

	r := utils.TestRouter()
	r.GET("/character/:id/phrases", GetAllPhrasesForCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPhrasesDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	mockRepo := &phrasesMockRepository{}
	phraseRepository = mockRepo

	mockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).Return([]Phrase{}, false, errors.New("DB error"))

	req := httptest.NewRequest(http.MethodGet, "/character/1/phrase", nil)

	r := utils.TestRouter()
	r.GET("/character/:id/phrase", GetAllPhrasesForCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetAllPhrasesForCharacterNotFound(t *testing.T) {
	t.Log("Phrases not found error should return Not Found")

	w := httptest.NewRecorder()

	mockRepo := &phrasesMockRepository{}
	phraseRepository = mockRepo

	mockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).Return([]Phrase{}, false, nil)

	req := httptest.NewRequest(http.MethodGet, "/character/1/phrases", nil)

	r := utils.TestRouter()
	r.GET("/character/:id/phrases", GetAllPhrasesForCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetAllPhrasesForCharacterFound(t *testing.T) {
	t.Log("Found phrases should return all phrases")

	w := httptest.NewRecorder()

	mockRepo := &phrasesMockRepository{}
	phraseRepository = mockRepo

	now := time.Now()
	phrases := []Phrase{
		NewPhrase(1, 1, nil, "miameeee", now, now),
		NewPhrase(2, 1, nil, "el tren de Ricardo Fort pasa una sola vez en la vida", now, now),
	}
	mockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).Return(phrases, true, nil)

	phResult := make([]PhraseResult, len(phrases))
	for i, phrase := range phrases {
		phResult[i] = PhraseResultFromPhrase(phrase)
	}

	req := httptest.NewRequest(http.MethodGet, "/character/1/phrases", nil)

	r := utils.TestRouter()
	r.GET("/character/:id/phrases", GetAllPhrasesForCharacter)
	r.ServeHTTP(w, req)

	actualResult := make(map[string][]PhraseResult)
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	resultPhrases := actualResult["results"]

	for i, phraseResult := range resultPhrases {
		assert.Equal(t, phResult[i].ID, phraseResult.ID)
		assert.Equal(t, phResult[i].Content, phraseResult.Content)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSavePhraseDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	mockRepo := &phrasesMockRepository{}
	phraseRepository = mockRepo

	mockRepo.On("Save", mock.Anything, mock.Anything).Return(Phrase{}, errors.New("DB error"))

	chCmd := NewPhraseCommand(1, "miameee")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPost, "/character/1/phrase", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character/:id/phrase", SaveNewPhrase)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSaveCharacterBadBodyFormat(t *testing.T) {
	t.Log("Bad body format should return Bad Request")

	w := httptest.NewRecorder()

	bodyStr := `{"wrong_field":"wrong_value"}`

	body := []byte(bodyStr)
	req := httptest.NewRequest(http.MethodPost, "/character/1/phrase", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character/:id/phrase", SaveNewPhrase)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSavePhraseOk(t *testing.T) {
	t.Log("Correct command should save and return Phrase")

	w := httptest.NewRecorder()

	mockRepo := &phrasesMockRepository{}
	phraseRepository = mockRepo

	now := time.Now()
	ph := NewPhrase(1, 1, nil, "miameee", now, now)
	mockRepo.On("Save", mock.Anything, mock.Anything).Return(ph, nil)

	phCmd := NewPhraseCommand(1, "miameee")
	body, _ := json.Marshal(phCmd)
	req := httptest.NewRequest(http.MethodPost, "/character/1/phrase", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character/:id/phrase", SaveNewPhrase)
	r.ServeHTTP(w, req)

	phResult := PhraseResultFromPhrase(ph)
	actualResult := PhraseResult{}
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	assert.Equal(t, phResult.ID, actualResult.ID)
	assert.Equal(t, phResult.Content, actualResult.Content)
	assert.Equal(t, http.StatusOK, w.Code)
}

type phrasesMockRepository struct {
	mock.Mock
}

func (repoMock *phrasesMockRepository) GetAllForCharacter(c *gin.Context, characterId int64) ([]Phrase, bool, error) {
	args := repoMock.Called(c, characterId)

	ch, ok := args.Get(0).([]Phrase)
	if !ok {
		panic(errors.New("mock error"))
	}

	found, ok := args.Get(1).(bool)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, found, args.Error(2)
}

func (repoMock *phrasesMockRepository) Save(c *gin.Context, phCmd PhraseCommand) (Phrase, error) {
	args := repoMock.Called(c, phCmd)

	ch, ok := args.Get(0).(Phrase)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, args.Error(1)
}
