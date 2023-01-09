package app

import (
	"github.com/sirupsen/logrus"
	"time"
)

type App struct {
	Server *fiber.App
}

func NewApp() *App {
	a := new(App)
	config := fiber.Config{
		AppName:      "Logger",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	a.Server = fiber.New(config)
	return a
}

func (a *App) Run(port string) {
	if err := a.Server.Listen(":" + port); err != nil {
		logrus.Fatal(err)
	}
}
