package main

import (
	"github.com/cherryReptile/WS-APP/internal/app"
	"github.com/cherryReptile/WS-APP/internal/middlewares"
	"github.com/cherryReptile/WS-APP/internal/sqlite"
	"github.com/cherryReptile/WS-AUTH/grpc/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	app := app.NewApp()
	app.Server.Use(logger.New())
	app.Server.Get("/test", func(ctx *fiber.Ctx) error {
		if err := sqlite.Create("test"); err != nil {
			logrus.Warning(err)
			return err
		}
		return ctx.JSON(map[string]string{"message": "database created successfully"})
	})

	//init client connection to grpc server
	conn, err := client.NewConn("auth:9000")

	if err != nil {
		logrus.Warning(err)
	}

	grpcClients := new(client.ServiceClients)
	grpcClients.Init(conn)

	home := app.Server.Group("/home", middlewares.CheckAuth(grpcClients.CheckAuth))
	home.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "ok"})
	})

	errCh := make(chan error)
	go app.Run("80", errCh)

	if err := <-errCh; err != nil {
		logrus.Fatal(err)
	}
}
