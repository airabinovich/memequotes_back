package character

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

func TestGetCharacterNonNumericIdShouldFail(t *testing.T) {
	t.Log("Calling with a non-numeric ID should return an error")

	w := httptest.NewRecorder()


	req := httptest.NewRequest(http.MethodGet, "/character/john", nil)

	r := utils.TestRouter()
	r.GET("/character/:id", GetCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetCharacterDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	mockRepo := &characterMockRepository{}
	characterRepository = mockRepo

	mockRepo.On("Get", mock.Anything, mock.Anything).Return(Character{}, false, errors.New("DB error"))

	req := httptest.NewRequest(http.MethodGet, "/character/1", nil)

	r := utils.TestRouter()
	r.GET("/character/:id", GetCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetCharacterNotFound(t *testing.T) {
	t.Log("Character not found error should return Not Found")

	w := httptest.NewRecorder()

	mockRepo := &characterMockRepository{}
	characterRepository = mockRepo

	mockRepo.On("Get", mock.Anything, mock.Anything).Return(Character{}, false, nil)

	req := httptest.NewRequest(http.MethodGet, "/character/1", nil)

	r := utils.TestRouter()
	r.GET("/character/:id", GetCharacter)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetCharacterFound(t *testing.T) {
	t.Log("Found character should return character")

	w := httptest.NewRecorder()

	mockRepo := &characterMockRepository{}
	characterRepository = mockRepo

	now := time.Now()
	ch := NewCharacter(1, "Comandante Fort", now, now)
	mockRepo.On("Get", mock.Anything, mock.Anything).Return(ch, true, nil)

	chResult := CharacterResultFromCharacter(ch)

	req := httptest.NewRequest(http.MethodGet, "/character/1", nil)

	r := utils.TestRouter()
	r.GET("/character/:id", GetCharacter)
	r.ServeHTTP(w, req)

	actualResult := CharacterResult{}
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	assert.Equal(t, chResult.ID, actualResult.ID)
	assert.Equal(t, chResult.Name, actualResult.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSaveCharacterDBFails(t *testing.T) {
	t.Log("DB error should return Internal Server Error")

	w := httptest.NewRecorder()

	mockRepo := &characterMockRepository{}
	characterRepository = mockRepo

	mockRepo.On("Save", mock.Anything, mock.Anything).Return(Character{}, errors.New("DB error"))

	chCmd := NewCharacterCommand("Comandante Fort")
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

	mockRepo := &characterMockRepository{}
	characterRepository = mockRepo

	now := time.Now()
	ch := NewCharacter(1, "Comandante Fort", now, now)
	mockRepo.On("Save", mock.Anything, mock.Anything).Return(ch, nil)

	chCmd := NewCharacterCommand("Comandante Fort")
	body, _ := json.Marshal(chCmd)
	req := httptest.NewRequest(http.MethodPost, "/character", bytes.NewBuffer(body))

	r := utils.TestRouter()
	r.POST("/character", SaveCharacter)
	r.ServeHTTP(w, req)

	chResult := CharacterResultFromCharacter(ch)
	actualResult := CharacterResult{}
	assert.NoError(t, json.NewDecoder(w.Result().Body).Decode(&actualResult))

	assert.Equal(t, chResult.ID, actualResult.ID)
	assert.Equal(t, chResult.Name, actualResult.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}

type characterMockRepository struct {
	mock.Mock
}

func (repoMock *characterMockRepository) Get(c *gin.Context, id int64) (Character, bool, error) {
	args := repoMock.Called(c, id)

	ch, ok := args.Get(0).(Character)
	if !ok {
		panic(errors.New("mock error"))
	}

	found, ok := args.Get(1).(bool)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, found, args.Error(2)
}

func (repoMock *characterMockRepository) Save(c *gin.Context, chCmd CharacterCommand) (Character, error) {
	args := repoMock.Called(c, chCmd)

	ch, ok := args.Get(0).(Character)
	if !ok {
		panic(errors.New("mock error"))
	}

	return ch, args.Error(1)
}
