package server

import (
	"go-api-fiber-starterkit/server/handler"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	server := fiber.New()

	handler.RegisterUser(server)

	return server
}