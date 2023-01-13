package server

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type HttpServer struct {
	Server *fiber.App
}

func NewHttpServer() *HttpServer {
	a := new(HttpServer)
	config := fiber.Config{
		AppName:      "Logger",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	a.Server = fiber.New(config)
	return a
}

func (a *HttpServer) Run(port string, errCh chan error) {
	if err := a.Server.Listen(":" + port); err != nil {
		errCh <- err
		return
	}
}
