package routers

import (
	"github.com/frostyjet/mythology-gallery-api/pkg/middleware"
	v1 "github.com/frostyjet/mythology-gallery-api/pkg/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORSMiddleware())

	// Handlers V1
	handlerV1 := v1.NewHandler()
	router.GET("/gods", handlerV1.GetAllGods)
	router.GET("/gods/:slug", handlerV1.GetGod)

	router.GET("/generate/gods", handlerV1.GenerateGods)
	router.GET("/generate/god-details", handlerV1.GenerateGodDetails)

	return router
}
