package handlers

import (
	"dust-api-service/internal/api"
	"dust-api-service/internal/db"
	"dust-api-service/internal/models"
	"dust-api-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
)

func VerifyAdminPermissions(username string) bool {
	userRole, _ := api.GetUserRole(username)
	return userRole == "admin"
}

func InitAdminHandlers(app *fiber.App) {
	apiGroup := app.Group("/api")
	admin := apiGroup.Group("/admin", func(c *fiber.Ctx) error { // middleware for /api/v1
		username := tokens.GetUsernameFromToken(c)
		if VerifyAdminPermissions(username) {
			return c.Next()
		} else {
			return c.SendStatus(403)
		}
	})

	admin.Put("/change_role", func(c *fiber.Ctx) error {
		userData := new(models.ChangeRoleRequest)
		if err := c.BodyParser(userData); err != nil {
			return err
		}
		err := db.UpdateRole(userData.Username, userData.Role)
		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		} else {
			return c.JSON(fiber.Map{"success": true})
		}
	})

}
