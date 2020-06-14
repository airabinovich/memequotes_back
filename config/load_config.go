package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-akka/configuration"
)

// Conf holds project main configuration
var Conf *configuration.Config

// Credentials holds all credentials
var Credentials *configuration.Config

// Config holds project main configuration
type Config struct {
	LogConfig LogConfig `json:"log"`
}

// LogConfig represents main configuration for logging
type LogConfig struct {
	FilePath string `json:"file_path"`
	Level    string `json:"level"`
	MaxAge   int    `json:"max_age"` // It's express in days. How many days we keep the logs
}

// LoadConfiguration loads main configuration given the fileName
func LoadConfiguration(fileName string) {
	path, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal("Configuration path error ", err)
	}
	Conf = configuration.LoadConfig(path)
}

// LoadCredentials loads credentials configuration
func LoadCredentials() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println("ERROR: could not get home directory")
		return
	}
	path := fmt.Sprintf("%s/credentials.conf", homedir)
	if _, err = os.Stat(path); os.IsNotExist(err) {
		log.Println("ERROR: could not find config file")
		return
	}
	Credentials = configuration.LoadConfig(path)
}
