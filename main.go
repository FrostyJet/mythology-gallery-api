package main

import (
	"github.com/frostyjet/mythology-gallery-api/routers"
)

func main() {
	routersInit := routers.InitRouter()

	routersInit.Run(":8282")
}
