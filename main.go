package main

import (
	"log"

	"github.com/frostyjet/mythology-gallery-api/models"
	"github.com/frostyjet/mythology-gallery-api/pkg/setting"
	"github.com/frostyjet/mythology-gallery-api/routers"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	setting.Setup()
	models.Setup()
}

func main() {
	routersInit := routers.InitRouter()

	routersInit.Run(":8282")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
