package main

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/handlers"
	"dust-api-service/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Dust API
// @version 1.0
// @description This is a API for dust server
// @contact.email fiber@swagger.io
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDatabase()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/docs", "./static/docs")

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/docs/swagger.yaml",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))

	handlers.InitAuthHandlers(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))

	handlers.InitAPI(app)
	handlers.InitSafetyHandlers(app)
	handlers.InitAdminHandlers(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.GracefulStartServer(app)
	}

}
