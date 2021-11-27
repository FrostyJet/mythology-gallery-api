package setting

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Environment struct{}

func (e *Environment) Setup() {
	path, err := e.GetConfigsPath()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (e *Environment) GetConfigsPath() (string, error) {
	path, err := filepath.Abs("../environments/.env")

	return path, err
}
