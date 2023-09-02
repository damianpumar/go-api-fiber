package handler

import (
	"go-api-fiber-starterkit/services/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(server *fiber.App){
	group :=server.Group("v1")

	group.Get("/users",  getAll)
}

func getAll(c *fiber.Ctx) error {
	userService := user.New()

	users, err := userService.GetUsers()

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(users)
}