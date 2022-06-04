package handlers

import (
	"dust-api-service/internal/api"
	"github.com/gofiber/fiber/v2"
)

func InitAuthHandlers(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		err := api.CreateUser()
		return c.SendString("Hello, World 👋!")
	})
}

func InitRestrictedAPI(app *fiber.App) {
	api := app.Group("/api")

}
