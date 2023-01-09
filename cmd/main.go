package main

import (
	"github.com/cherryReptile/WS-APP/internal/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	app := app.NewApp()
	app.Server.Use(logger.New())
	app.Server.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("test message")
	})

	errCh := make(chan error)
	go app.Run("80", errCh)

	if err := <-errCh; err != nil {
		logrus.Fatal(err)
	}
}
