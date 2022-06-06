package handlers

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/models"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
)

func InitAdminHandlers(app *fiber.App) {
	apiGroup := app.Group("/api")
	admin := apiGroup.Group("/admin", func(c *fiber.Ctx) error { // middleware for /api/v1
		username := tokens.GetUsernameFromToken(c)
		if VerifyPermissions(username, "admin") {
			return c.Next()
		} else {
			return c.SendStatus(403)
		}
	})

	admin.Get("change_role", func(c *fiber.Ctx) error {
		userData := new(models.ChangeRoleRequest)
		err := db.UpdateRole(userData.Username, userData.Role)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		} else {
			return c.JSON(fiber.Map{"success": true})
		}
	})
}
