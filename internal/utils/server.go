package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func buildConnectionURL() string {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func StartServer(app *fiber.App) {
	connURL := buildConnectionURL()

	// Run server
	if err := app.Listen(connURL); err != nil {
		log.Error("Server is not running! Reason: %v", err)
	}
}

func GracefulStartServer(app *fiber.App) {
	connURL := buildConnectionURL()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Info("Received SIGINT. Shutting down.")
		if err := app.Shutdown(); err != nil {
			log.Error("Server is not shutting down! Reason: %v", err)
		}
	}()

	// Run server
	fmt.Println(connURL)
	if err := app.Listen(connURL); err != nil {
		log.Error("Server is not running! Reason: %v", err)
	}
}
