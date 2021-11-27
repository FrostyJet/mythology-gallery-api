package v1

import (
	"net/http"

	"github.com/frostyjet/mythology-gallery-api/pkg/models"
	"github.com/gin-gonic/gin"
)

func (g *Handler) GetAllGods(c *gin.Context) {
	data, _ := models.GetGodsList()

	c.JSON(200, gin.H{
		"gods": data,
	})
}

func (g *Handler) GetGod(c *gin.Context) {
	slug := c.Param("slug")

	godsModel := models.God{}
	data, err := godsModel.GetBySlug(slug)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "God not found",
			"slug":    slug,
		})
		return
	}

	c.JSON(200, gin.H{
		"god": data,
	})
}
