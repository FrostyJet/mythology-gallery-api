package routers

import (
	"github.com/frostyjet/mythology-gallery-api/models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/gods", func(c *gin.Context) {
		data, _ := models.GetGodsList()

		c.JSON(200, gin.H{
			"gods": data,
		})
	})

	return router
}
