package main

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/handlers"
	"dust-api-service/internal/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDatabase()
	app := fiber.New()

	handlers.InitAuthHandlers(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))

	handlers.InitAPI(app)

	handlers.InitSafetyHandlers(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.GracefulStartServer(app)
	}

}
