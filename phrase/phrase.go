package phrase

import (
	"github.com/airabinovich/memequotes_back/character"
	"github.com/airabinovich/memequotes_back/utils"
	"time"
)

// PhraseResult is the type to be shown in the API for a Phrase
type PhraseResult struct {
	ID          int64              `json:"id"`
	CharacterId int64              `json:"character_id"`
	Content     string             `json:"content"`
	DateCreated *utils.ISO8601Time `json:"date_created"`
	LastUpdated *utils.ISO8601Time `json:"last_updated"`
}

// NewPhraseResult is a constructor for PhraseResult
func NewPhraseResult(ID int64, characterId int64, content string, dateCreacted *utils.ISO8601Time, lastUpdated *utils.ISO8601Time) PhraseResult {
	return PhraseResult{
		ID:          ID,
		CharacterId: characterId,
		Content:     content,
		DateCreated: dateCreacted,
		LastUpdated: lastUpdated,
	}
}

// PhraseResultFromPhrase creates a PhraseResult from a Phrase
func PhraseResultFromPhrase(phrase Phrase) PhraseResult {
	dateCreated := utils.ISO8601Time(phrase.DateCreated)
	lastUpdated := utils.ISO8601Time(phrase.LastUpdated)
	return PhraseResult{
		ID:          phrase.ID,
		CharacterId: phrase.CharacterId,
		Content:     phrase.Content,
		DateCreated: &dateCreated,
		LastUpdated: &lastUpdated,
	}
}

// Phrase represent a phrase from one character
type Phrase struct {
	ID          int64 `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterId int64
	Character   *character.Character `gorm:"foreignkey:CharacterId"`
	Content     string
	DateCreated time.Time `gorm:"column:date_created;type:datetime;not null"`
	LastUpdated time.Time `gorm:"column:last_updated;type:datetime;not null"`
}

// NewPhrase is a constructor for Phrase
func NewPhrase(ID int64, characterId int64, character *character.Character, content string, dateCreacted time.Time, lastUpdated time.Time) Phrase {
	return Phrase{
		ID:          ID,
		CharacterId: characterId,
		Character:   character,
		Content:     content,
		DateCreated: dateCreacted,
		LastUpdated: lastUpdated,
	}
}

// PhraseCommand contains the info to create a phrase
type PhraseCommand struct {
	CharacterId int64  `json:"character_id" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

// NewPhraseCommand is a constructor for PhraseCommand
func NewPhraseCommand(characterId int64, content string) PhraseCommand {
	return PhraseCommand{
		CharacterId: characterId,
		Content: content,
	}
}
