package handlers

import (
	"dust-api-service/internal/api"
	"dust-api-service/internal/db"
	"dust-api-service/internal/models"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func InitAPI(app *fiber.App) {
	apiGroup := app.Group("/api")

	character := apiGroup.Group("/character")

	character.Put("/add", func(c *fiber.Ctx) error {
		requestJson := new(models.CharacterRequest)
		username := tokens.GetUsernameFromToken(c)
		if err := c.BodyParser(requestJson); err != nil {
			return err
		}
		character := models.CharacterData{
			Username:         username,
			CharacterRequest: *requestJson,
		}
		err := api.CreateCharacter(character)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		} else {
			return c.JSON(fiber.Map{"success": true})
		}
	})
	character.Get("/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")
		userCharacters := db.GetAllUserCharacters(username)
		return c.JSON(userCharacters)
	})
	character.Put("/edit", func(c *fiber.Ctx) error {
		requestJson := new(models.CharacterRequest)
		username := tokens.GetUsernameFromToken(c)
		if err := c.BodyParser(requestJson); err != nil {
			return err
		}
		err := db.UpdateCharacter(username, requestJson.Title, requestJson.Description)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		} else {
			return c.JSON(fiber.Map{"success": true})
		}
	})
	character.Delete("/delete", func(c *fiber.Ctx) error {
		username := tokens.GetUsernameFromToken(c)
		titleRequest := struct {
			Title string `json:"title"`
		}{}

		if err := c.BodyParser(&titleRequest); err != nil {
			return c.JSON(fiber.Map{"error": "invalid request"})
		}

		err := api.RemoveCharacter(username, titleRequest.Title)
		if err != nil {
			log.Error(err)
			return c.JSON(fiber.Map{"error": true})
		}
		return c.JSON(fiber.Map{"success": true})
	})
}
