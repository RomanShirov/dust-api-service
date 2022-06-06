package handlers

import (
	"dust-api-service/internal/api"
	"dust-api-service/internal/db"
	"dust-api-service/internal/models"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
)

func VerifyModeratorPermissions(username string) bool {
	userRole, _ := api.GetUserRole(username)
	return userRole == "admin" || userRole == "moderator"
}

func InitSafetyHandlers(app *fiber.App) {
	apiGroup := app.Group("/api")
	safety := apiGroup.Group("/safety", func(c *fiber.Ctx) error { // middleware for /api/v1
		username := tokens.GetUsernameFromToken(c)
		if VerifyModeratorPermissions(username) {
			return c.Next()
		} else {
			return c.SendStatus(403)
		}
	})

	safety.Get("/get_all_users", func(c *fiber.Ctx) error {
		users := db.GetAllUsers()
		return c.JSON(users)
	})

	safety.Get("/get_all_characters", func(c *fiber.Ctx) error {
		users := db.GetAllCharacters()
		return c.JSON(users)
	})

	safety.Delete("/remove_character", func(c *fiber.Ctx) error {
		request := new(models.RemoveCharacterRequest)
		if err := c.BodyParser(request); err != nil {
			return err
		}
		err := api.RemoveCharacter(request.Username, request.Title)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		} else {
			return c.JSON(fiber.Map{"success": true})
		}
	})
}
