package handlers

import (
	"dust-api-service/internal/api"
	"dust-api-service/internal/db"
	"dust-api-service/internal/tokens"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	//log "github.com/sirupsen/logrus"
)

func InitAuthHandlers(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/token", func(c *fiber.Ctx) error {
		fmt.Println(os.Getenv("JWT_SECRET_KEY"))
		fmt.Println("SK:", []byte(os.Getenv("JWT_SECRET_KEY")))
		return c.SendString("Hello")
	})
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

}
