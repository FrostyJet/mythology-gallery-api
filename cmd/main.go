package main

import (
	"github.com/frostyjet/mythology-gallery-api/models"
	"github.com/frostyjet/mythology-gallery-api/pkg/routers"
	"github.com/frostyjet/mythology-gallery-api/pkg/setting"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	routersInit := routers.InitRouter()

	routersInit.Run(":" + setting.ServerSetting.Port)
}
