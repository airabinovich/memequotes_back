package character

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/airabinovich/memequotes_back/model"
	"github.com/airabinovich/memequotes_back/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	characterMockRepo characterMockRepository
	phraseMockRepo    phrasesMockRepository
)

func TestGetCharacterNonNumericIdShouldFail(t *testing.T) {
	t.Log("Calling with a non-numeric ID should return an error")

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/character/john", nil)

	r := utils.TestRouter()
	r.GET("/character/:character-id", GetCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetCharacterDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	resetMocks()

	characterMockRepo.On("Get", mock.Anything, mock.Anything).Return(model.Character{}, false, errors.New("DB error"))

	req := httptest.NewRequest(http.MethodGet, "/character/1", nil)

	r := utils.TestRouter()
	r.GET("/character/:character-id", GetCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetCharacterNotFound(t *testing.T) {
	t.Log("Character not found error should return Not Found")

	w := httptest.NewRecorder()

	resetMocks()

	characterMockRepo.On("Get", mock.Anything, mock.Anything).Return(model.Character{}, false, nil)

	req := httptest.NewRequest(http.MethodGet, "/character/1", nil)

	r := utils.TestRouter()
	r.GET("/character/:character-id", GetCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetCharacterFound(t *testing.T) {
	t.Log("Found character should return character")

	w := httptest.NewRecorder()

	resetMocks()

	now := time.Now()
	ch := model.NewCharacter(1, "Comandante Fort", now, now)
	characterMockRepo.On("Get", mock.Anything, mock.Anything).Return(ch, true, nil)

	chResult := model.CharacterResultFromCharacter(ch)

	req := httptest.NewRequest(http.MethodGet, "/character/1", nil)

	r := utils.TestRouter()
	r.GET("/character/:character-id", GetCharacter)
	r.ServeHTTP(w, req)

	actualResult := model.CharacterResult{}
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	assert.Equal(t, chResult.ID, actualResult.ID)
	assert.Equal(t, chResult.Name, actualResult.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllCharactersDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	resetMocks()

	characterMockRepo.On("GetAll", mock.Anything, mock.Anything).Return([]model.Character{}, errors.New("DB error"))

	req := httptest.NewRequest(http.MethodGet, "/characters", nil)

	r := utils.TestRouter()
	r.GET("/characters", GetAllCharacters)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetAllCharactersOK(t *testing.T) {
	t.Log("Found all characters should return characters")

	w := httptest.NewRecorder()

	resetMocks()

	now := time.Now()
	chs := []model.Character{
		model.NewCharacter(1, "Comandante Fort", now, now),
		model.NewCharacter(1, "Guillermo Franchella", now, now),
	}
	characterMockRepo.On("GetAll", mock.Anything, mock.Anything).Return(chs, nil)

	chResult := make([]model.CharacterResult, len(chs))
	for i, ch := range chs {
		chResult[i] = model.CharacterResultFromCharacter(ch)
	}

	req := httptest.NewRequest(http.MethodGet, "/characters", nil)

	r := utils.TestRouter()
	r.GET("/characters", GetAllCharacters)
	r.ServeHTTP(w, req)

	actualResult := make(map[string][]model.CharacterResult)
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	for i, result := range actualResult["results"] {
		assert.Equal(t, chResult[i].ID, result.ID)
		assert.Equal(t, chResult[i].Name, result.Name)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSaveCharacterDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	resetMocks()

	characterMockRepo.On("Save", mock.Anything, mock.Anything).Return(model.Character{}, errors.New("DB error"))

	chCmd := model.NewCharacterCommand("Comandante Fort")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPost, "/character", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character", SaveCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSaveCharacterBadBodyFormat(t *testing.T) {
	t.Log("Bad body format should return Bad Request")

	w := httptest.NewRecorder()

	bodyStr := `{"wrong_field":"wrong_value"}`
	body := []byte(bodyStr)
	req := httptest.NewRequest(http.MethodPost, "/character", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character", SaveCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSaveCharacterOk(t *testing.T) {
	t.Log("Correct command should save and return Character")

	w := httptest.NewRecorder()

	resetMocks()

	now := time.Now()
	ch := model.NewCharacter(1, "Comandante Fort", now, now)
	characterMockRepo.On("Save", mock.Anything, mock.Anything).Return(ch, nil)

	chCmd := model.NewCharacterCommand("Comandante Fort")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPost, "/character", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character", SaveCharacter)
	r.ServeHTTP(w, req)

	chResult := model.CharacterResultFromCharacter(ch)
	actualResult := model.CharacterResult{}
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	assert.Equal(t, chResult.ID, actualResult.ID)
	assert.Equal(t, chResult.Name, actualResult.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCharacterNonNumericIdShouldFail(t *testing.T) {
	t.Log("Calling with a non-numeric ID should return an error")

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPatch, "/character/john", nil)

	r := utils.TestRouter()
	r.PATCH("/character/:character-id", UpdateCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateCharacterWrongBodyShouldFail(t *testing.T) {
	t.Log("Calling with a non-numeric ID should return an error")

	w := httptest.NewRecorder()

	bodyStr := `{"wrong_field":"wrong_value"}`
	body := []byte(bodyStr)
	req := httptest.NewRequest(http.MethodPatch, "/character/1", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.PATCH("/character/:character-id", UpdateCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateCharacterDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	resetMocks()

	characterMockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(model.Character{}, false, errors.New("DB error"))

	chCmd := model.NewCharacterCommand("Comandante Fort")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPatch, "/character/1", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.PATCH("/character/:character-id", UpdateCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateCharacterNotFound(t *testing.T) {
	t.Log("Character not found error should return Not Found")

	w := httptest.NewRecorder()

	resetMocks()

	characterMockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(model.Character{}, false, nil)

	chCmd := model.NewCharacterCommand("Comandante Fort")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPatch, "/character/1", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.PATCH("/character/:character-id", UpdateCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateCharacterOK(t *testing.T) {
	t.Log("Update character should update and return character")

	w := httptest.NewRecorder()

	resetMocks()

	now := time.Now()
	ch := model.NewCharacter(1, "Comandante Fort", now, now)
	characterMockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(ch, true, nil)

	chCmd := model.NewCharacterCommand("Comandante Fort")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPost, "/character/1", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character/:character-id", UpdateCharacter)
	r.ServeHTTP(w, req)

	chResult := model.CharacterResultFromCharacter(ch)
	actualResult := model.CharacterResult{}
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	assert.Equal(t, chResult.ID, actualResult.ID)
	assert.Equal(t, chResult.Name, actualResult.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCharacterNonNumericId(t *testing.T) {
	t.Log("Calling with a non-numeric ID should return an error")

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/character/john", nil)

	r := utils.TestRouter()
	r.DELETE("/character/:character-id", DeleteCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteCharacterPhrasesDBError(t *testing.T) {
	t.Log("Phrases repository fails should return an error")

	resetMocks()

	phraseMockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).
		Return([]model.Phrase{}, false, errors.New("DB error"))

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/character/1", nil)

	r := utils.TestRouter()
	r.DELETE("/character/:character-id", DeleteCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteCharacterCannotDeletePhrases(t *testing.T) {
	t.Log("Phrases repository fails deleting phrases should return an error")

	resetMocks()

	now := time.Now()
	phrases := []model.Phrase{
		{
			ID:          1,
			CharacterId: 1,
			Character:   nil,
			Content:     "Jojoojo",
			DateCreated: now,
			LastUpdated: now,
		},
	}
	phraseMockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).
		Return(phrases, true, nil)

	phraseMockRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("cannot delete phrase"))

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/character/1", nil)

	r := utils.TestRouter()
	r.DELETE("/character/:character-id", DeleteCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteCharacterDBError(t *testing.T) {
	t.Log("Character repository fails should return an error")

	resetMocks()

	now := time.Now()
	phrases := []model.Phrase{
		{
			ID:          1,
			CharacterId: 1,
			Character:   nil,
			Content:     "Jojoojo",
			DateCreated: now,
			LastUpdated: now,
		},
	}
	phraseMockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).
		Return(phrases, true, nil)

	phraseMockRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	characterMockRepo.On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("DB error"))

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/character/1", nil)

	r := utils.TestRouter()
	r.DELETE("/character/:character-id", DeleteCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteCharacterOK(t *testing.T) {
	t.Log("Delete Character should return Gone")

	resetMocks()

	now := time.Now()
	phrases := []model.Phrase{
		{
			ID:          1,
			CharacterId: 1,
			Character:   nil,
			Content:     "Jojoojo",
			DateCreated: now,
			LastUpdated: now,
		},
	}
	phraseMockRepo.On("GetAllForCharacter", mock.Anything, mock.Anything).
		Return(phrases, true, nil)

	phraseMockRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	characterMockRepo.On("Delete", mock.Anything, mock.Anything).
		Return(nil)

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodDelete, "/character/1", nil)

	r := utils.TestRouter()
	r.DELETE("/character/:character-id", DeleteCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusGone, w.Code)
	characterMockRepo.AssertCalled(t, "Delete", mock.Anything, mock.Anything)
}

func resetMocks() {
	characterMockRepo = characterMockRepository{}
	phraseMockRepo = phrasesMockRepository{}
	characterRepository = &characterMockRepo
	phraseRepository = &phraseMockRepo
}

type characterMockRepository struct {
	mock.Mock
}

func (repoMock *characterMockRepository) Get(c *gin.Context, id int64) (model.Character, bool, error) {
	args := repoMock.Called(c, id)

	ch, ok := args.Get(0).(model.Character)
	if !ok {
		panic(errors.New("mock error"))
	}

	found, ok := args.Get(1).(bool)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, found, args.Error(2)
}

func (repoMock *characterMockRepository) Save(c *gin.Context, chCmd model.CharacterCommand) (model.Character, error) {
	args := repoMock.Called(c, chCmd)

	ch, ok := args.Get(0).(model.Character)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, args.Error(1)
}

func (repoMock *characterMockRepository) Update(c *gin.Context, id int64, chCmd model.CharacterCommand) (model.Character, bool, error) {
	args := repoMock.Called(c, id, chCmd)

	ch, ok := args.Get(0).(model.Character)
	if !ok {
		panic(errors.New("mock error"))
	}

	found, ok := args.Get(1).(bool)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, found, args.Error(2)
}

func (repoMock *characterMockRepository) GetAll(c *gin.Context) ([]model.Character, error) {
	args := repoMock.Called(c)

	ch, ok := args.Get(0).([]model.Character)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, args.Error(1)
}

func (repoMock *characterMockRepository) Delete(c *gin.Context, id int64) error {
	args := repoMock.Called(c, id)

	return args.Error(0)
}

type phrasesMockRepository struct {
	mock.Mock
}

func (repoMock *phrasesMockRepository) Get(c *gin.Context, characterId int64, id int64) (model.Phrase, bool, error) {
	args := repoMock.Called(c, characterId, id)

	ph, ok := args.Get(0).(model.Phrase)
	if !ok {
		panic(errors.New("mock error"))
	}

	found, ok := args.Get(1).(bool)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ph, found, args.Error(2)
}

func (repoMock *phrasesMockRepository) GetAllForCharacter(c *gin.Context, characterId int64) ([]model.Phrase, bool, error) {
	args := repoMock.Called(c, characterId)

	ph, ok := args.Get(0).([]model.Phrase)
	if !ok {
		panic(errors.New("mock error"))
	}

	found, ok := args.Get(1).(bool)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ph, found, args.Error(2)
}

func (repoMock *phrasesMockRepository) Save(c *gin.Context, phCmd model.PhraseCommand) (model.Phrase, error) {
	args := repoMock.Called(c, phCmd)

	ph, ok := args.Get(0).(model.Phrase)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ph, args.Error(1)
}

func (repoMock *phrasesMockRepository) Delete(c *gin.Context, characterId int64, id int64) error {
	args := repoMock.Called(c, characterId, id)
	return args.Error(0)
}
