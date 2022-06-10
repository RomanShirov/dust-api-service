package handlers

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
)

func InitAuthHandlers(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiGroup.Post("/register", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		if db.CheckUserExists(username, password) {
			_ = c.SendStatus(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{"error": "user already exists"})
		}
		token, err := db.CreateUser(username, password)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		} else {
			return c.JSON(fiber.Map{
				"success": true,
				"token":   token})
		}
	})

	apiGroup.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		userIsValid := db.ValidateUser(username, password)
		if userIsValid {
			token := tokens.GenerateUserToken(username)
			return c.JSON(fiber.Map{"success": true, "token": token})
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	})
}
