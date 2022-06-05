package handlers

import (
	"dust-api-service/internal/api"
	"dust-api-service/internal/db"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
)

func InitAuthHandlers(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiGroup.Post("/register", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		token, err := api.CreateUser(username, password)
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

func InitRestrictedAPI(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiGroup.Get("/test", func(c *fiber.Ctx) error {
		name := tokens.GetUsernameFromToken(c)
		return c.SendString(name)
	})
	apiGroup.Get("/getAllUsers", func(c *fiber.Ctx) error {
		users := db.GetAllUsers()
		return c.JSON(users)
	})

}
