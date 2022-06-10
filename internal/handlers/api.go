package handlers

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/models"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type CharacterTitle struct {
	Title string `json:"title"`
}

func InitAPI(app *fiber.App) {

	apiGroup := app.Group("/api")

	character := apiGroup.Group("/character")

	character.Post("/", func(c *fiber.Ctx) error {
		requestJson := new(models.CharacterRequest)
		username := tokens.GetUsernameFromToken(c)
		if err := c.BodyParser(requestJson); err != nil {
			return err
		}
		character := models.CharacterData{
			Username:         username,
			CharacterRequest: *requestJson,
		}
		err := db.CreateCharacter(character)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		}
		return c.JSON(fiber.Map{"success": true})
	})

	character.Get("/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")
		userCharacters := db.GetAllUserCharacters(username)
		return c.JSON(userCharacters)
	})

	character.Put("/", func(c *fiber.Ctx) error {
		requestJson := new(models.CharacterRequest)
		username := tokens.GetUsernameFromToken(c)
		if err := c.BodyParser(requestJson); err != nil {
			return err
		}
		err := db.UpdateCharacter(username, requestJson.Title, requestJson.Description)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		}
		return c.JSON(fiber.Map{"success": true})
	})

	character.Delete("/", func(c *fiber.Ctx) error {
		username := tokens.GetUsernameFromToken(c)

		titleRequest := new(CharacterTitle)
		if err := c.BodyParser(&titleRequest); err != nil {
			return c.JSON(fiber.Map{"error": "invalid request"})
		}

		err := db.RemoveCharacter(username, titleRequest.Title)
		if err != nil {
			log.Error(err)
			return c.JSON(fiber.Map{"error": true})
		}
		return c.JSON(fiber.Map{"success": true})
	})
}
