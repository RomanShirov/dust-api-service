package handlers

import (
	"dust-api-service/internal/api"
	"github.com/gofiber/fiber/v2"
	//log "github.com/sirupsen/logrus"
)

func InitAuthHandlers(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		token, err := api.CreateUser(username, password)
		if err != nil {
			return c.SendStatus(403)
		}
		return c.JSON(fiber.Map{
			"success": true,
			"token":   token})
	})

}

func InitRestrictedAPI(app *fiber.App) {
	//api := app.Group("/api")

}
