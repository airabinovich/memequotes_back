package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/airabinovich/memequotes_back/character"
	"github.com/airabinovich/memequotes_back/config"
	"github.com/airabinovich/memequotes_back/database"
	"github.com/airabinovich/memequotes_back/phrase"
	"github.com/airabinovich/memequotes_back/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

type commandFlags struct {
	credentialsFile string
}

func parseFlags() (commandFlags, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println("ERROR: could not get home directory")
		return commandFlags{}, errors.New("cannot get homedir")
	}
	defaultCredentialsFile := fmt.Sprintf("%s/credentials.conf", homedir)

	var credentialsFile string
	flag.StringVar(&credentialsFile, "credentials", defaultCredentialsFile, "The environment in which the application is running")
	flag.Parse()
	return commandFlags{
		credentialsFile: credentialsFile,
	}, nil
}

func main() {

	flags, _ := parseFlags()
	config.LoadCredentials(flags.credentialsFile)

	err := database.Initialize()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			panic(err)
		}
	}()

	character.Initialize()
	phrase.Initialize()

	engine := router.Route()
	if err := engine.Run(":9000"); err != nil {
		println("Backend service could not be started")
		panic(err)
	}
}
