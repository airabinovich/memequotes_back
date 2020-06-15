package character

import (
	"github.com/airabinovich/memequotes_back/utils"
	"time"
)

// CharacterResult is the type to be shown in the API for a Character
type CharacterResult struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	DateCreated *utils.ISO8601Time `json:"date_created"`
	LastUpdated *utils.ISO8601Time `json:"last_updated"`
}

// NewCharacterResult is a constructor for CharacterResult
func NewCharacterResult(id int64, name string, dateCreated *utils.ISO8601Time, lastUpdated *utils.ISO8601Time) CharacterResult {
	return CharacterResult{
		ID:          id,
		Name:        name,
		DateCreated: dateCreated,
		LastUpdated: lastUpdated,
	}
}

// CharacterResultFromCharacter creates a CharacterResult from a Character
func CharacterResultFromCharacter(ch Character) CharacterResult {
	dateCreated := utils.ISO8601Time(ch.DateCreated)
	lastUpdated := utils.ISO8601Time(ch.LastUpdated)
	return CharacterResult{
		ID:          ch.ID,
		Name:        ch.Name,
		DateCreated: &dateCreated,
		LastUpdated: &lastUpdated,
	}
}

// Character represents a character that may own prases
type Character struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string    `gorm:"unique"`
	DateCreated time.Time `gorm:"column:date_created;type:datetime;not null"`
	LastUpdated time.Time `gorm:"column:last_updated;type:datetime;not null"`
}

// NewCharacter is a constructor for Character
func NewCharacter(id int64, name string, dateCreated time.Time, lastUpdated time.Time) Character {
	return Character{
		ID:          id,
		Name:        name,
		DateCreated: dateCreated,
		LastUpdated: lastUpdated,
	}
}

//CharacterCommand contains the info to create a Character
type CharacterCommand struct {
	Name string `json:"name"`
}

//NewCharacterCommand is a constructor for CharacterCommand
func NewCharacterCommand(name string) CharacterCommand {
	return CharacterCommand{Name: name}
}
