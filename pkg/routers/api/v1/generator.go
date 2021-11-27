package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/frostyjet/mythology-gallery-api/pkg/models"
	"github.com/gin-gonic/gin"
)

type apiGodAbility struct {
	Id          int
	Summary     string
	URL         string
	Description struct {
		ItemDescription struct {
			Description string
		} `json:"ItemDescription"`
	}
}

type apiGodDetails struct {
	Id              int
	Slug            string
	Api_information struct {
		Ability_1 apiGodAbility
		Ability_2 apiGodAbility
		Ability_3 apiGodAbility
		Ability_4 apiGodAbility
		Ability_5 apiGodAbility
		Lore      string
	}
	Skins []struct {
		Name  string
		Type  string
		Image string
	}
	Acf struct {
		God_header_image string
	}
}

func (h *Handler) GenerateGods(c *gin.Context) {
	const apiUrl = "https://cms.smitegame.com/wp-json/smite-api/all-gods/1"

	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()
	contens, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	var data []models.God
	json.Unmarshal(contens, &data)

	models.InsertGodsList(&data)

	c.JSON(http.StatusOK, gin.H{
		"Inserted/Updated rows": len(data),
	})
}

func (h *Handler) GenerateGodDetails(c *gin.Context) {
	id := c.Query("id")
	amount := c.Query("amount")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Parameter 'id' must be valid integer",
		})
		return
	}

	amountNum, err := strconv.Atoi(amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Parameter 'amount' must be valid integer",
		})
		return
	}

	gods, err := models.SelectGodsFromWithLimit(idInt, amountNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	statuses := []string{}

	for _, god := range gods {
		slug := strings.ToLower(god.Name)
		slug = strings.ReplaceAll(slug, " ", "-")

		resp, err := http.Get("https://cms.smitegame.com/wp-json/wp/v2/gods?slug=" + slug + "&lang_id=1")
		if err != nil {
			statuses = append(statuses, god.Name+" - "+err.Error())
			continue
		}

		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			statuses = append(statuses, god.Name+" - "+err.Error())
			continue
		}

		data := make([]apiGodDetails, 0)
		json.Unmarshal(contents, &data)

		// Extract abilities into array for easy access
		abilitiesToInsert := []models.Ability{}
		abilities := []apiGodAbility{
			data[0].Api_information.Ability_1,
			data[0].Api_information.Ability_2,
			data[0].Api_information.Ability_3,
			data[0].Api_information.Ability_4,
			data[0].Api_information.Ability_5,
		}

		// Create models from abilities and add GodID relation
		for _, ability := range abilities {
			abilitiesToInsert = append(abilitiesToInsert, models.Ability{
				Id:          ability.Id,
				Summary:     ability.Summary,
				Description: ability.Description.ItemDescription.Description,
				URL:         ability.URL,
				GodID:       uint(god.Id),
			})
		}

		// Insert models into database
		abilityModel := models.Ability{}
		err = abilityModel.InsertMany(&abilitiesToInsert)

		if err == nil {
			statuses = append(statuses, god.Name+" - Abilities inserted!")
		}

		// Create models from skins with relation to GodID
		skinsToInsert := []models.Skin{}
		for _, skin := range data[0].Skins {
			skinsToInsert = append(skinsToInsert, models.Skin{
				Name:  skin.Name,
				Type:  skin.Type,
				Image: skin.Image,
				GodID: uint(god.Id),
			})
		}

		// Insert skins into database
		skinModel := models.Skin{}
		err = skinModel.InsertMany(&skinsToInsert)

		if err == nil {
			statuses = append(statuses, god.Name+" - Skins inserted!")
		}

		// Update god missing information
		err = god.Update(&models.God{
			Slug:        data[0].Slug,
			Lore:        data[0].Api_information.Lore,
			HeaderImage: data[0].Acf.God_header_image,
		})

		if err == nil {
			statuses = append(statuses, god.Name+" - Lore, slug and header-image updated!")
		}
	}

	if err != nil {
		statuses = append(statuses, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       idInt,
		"statuses": statuses,
	})
}
